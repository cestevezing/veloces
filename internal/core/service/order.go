package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/dto/response"
	"github.com/cestevezing/veloces/internal/core/model"
	"github.com/cestevezing/veloces/internal/core/port/repository"
	"github.com/cestevezing/veloces/internal/core/port/service"
)

type OrderImpl struct {
	repository        repository.IOrder
	productRepository repository.IProduct
}

func NewOrderService(repository repository.IOrder, productRepository repository.IProduct) service.IOrder {
	return &OrderImpl{
		repository:        repository,
		productRepository: productRepository,
	}
}

func (s *OrderImpl) Create(ctx context.Context, order *requests.OrderCreate) (*model.Order, error) {

	tx := s.repository.GetDB(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	newOrder := &model.Order{
		CustomerName: order.CustomerName,
	}

	for i, item := range order.Items {
		product, err := s.productRepository.FindOne(ctx, map[string]any{"id": item.ProductID}, tx)
		if err != nil {
			return nil, err
		}
		if product.Stock < item.Quantity {
			tx.Rollback()
			return nil, errors.New(fmt.Sprint("Insufficient stock for product ID: ", item.ProductID))
		}
		newOrder.Items = append(newOrder.Items, model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			OrderID:   newOrder.ID,
			Subtotal:  product.Price * float64(item.Quantity),
		})
		newOrder.TotalAmount += newOrder.Items[i].Subtotal
	}

	newOrder, err := s.repository.Create(ctx, newOrder)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range newOrder.Items {
		err = s.productRepository.UpdateStock(ctx, int(item.ProductID), item.Quantity, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return newOrder, nil
}

func (s *OrderImpl) GetByID(ctx context.Context, id int) (*response.Order, error) {
	orderResponse := &response.Order{}
	order, err := s.repository.FindOne(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}
	orderResponse.ID = order.ID
	orderResponse.CustomerName = order.CustomerName
	orderResponse.TotalAmount = order.TotalAmount
	orderResponse.Items = make([]response.Items, len(order.Items))
	for i, item := range order.Items {
		orderResponse.Items[i].ID = item.ProductID
		orderResponse.Items[i].Quantity = item.Quantity
		orderResponse.Items[i].Subtotal = item.Subtotal
		product, err := s.productRepository.FindOne(ctx, map[string]any{"id": item.ProductID})
		if err != nil {
			return nil, err
		}
		orderResponse.Items[i].Product.ID = product.ID
		orderResponse.Items[i].Product.Name = product.Name
		orderResponse.Items[i].Product.Price = product.Price
	}
	return orderResponse, nil
}
