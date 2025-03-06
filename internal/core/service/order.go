package service

import (
	"errors"
	"fmt"

	"github.com/cestevezing/veloces/internal/core/dto/requests"
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

func (s *OrderImpl) Create(order *requests.OrderCreate) (*model.Order, error) {
	newOrder := &model.Order{
		CustomerName: order.CustomerName,
	}

	productsUpdate := make([]model.Product, len(order.Items))
	for i, item := range order.Items {
		product, err := s.productRepository.FindOne(map[string]any{"id": item.ProductID})
		if err != nil {
			return nil, err
		}
		if product.Stock < item.Quantity {
			return nil, errors.New(fmt.Sprint("Insufficient stock for product ID: ", item.ProductID))
		}
		newOrder.Items = append(newOrder.Items, model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			OrderID:   newOrder.ID,
			Subtotal:  product.Price * float64(item.Quantity),
		})
		newOrder.TotalAmount += newOrder.Items[i].Subtotal
		product.Stock -= item.Quantity
		productsUpdate = append(productsUpdate, *product)
	}

	newOrder, err := s.repository.Create(newOrder)
	if err != nil {
		return nil, err
	}

	for _, item := range productsUpdate {
		_, err = s.productRepository.Update(&item)
		if err != nil {
			return nil, err
		}
	}

	return newOrder, nil
}

func (s *OrderImpl) GetByID(id int) (*model.Order, error) {
	return s.repository.FindOne(map[string]any{"id": id})
}
