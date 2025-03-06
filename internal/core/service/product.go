package service

import (
	"context"

	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/model"
	"github.com/cestevezing/veloces/internal/core/port/repository"
	"github.com/cestevezing/veloces/internal/core/port/service"
)

type ProductImpl struct {
	repository repository.IProduct
}

func NewProductService(repository repository.IProduct) service.IProduct {
	return &ProductImpl{repository: repository}
}

func (p *ProductImpl) GetAll(ctx context.Context) ([]*model.Product, error) {
	return p.repository.Find(ctx)
}

func (p *ProductImpl) GetByID(ctx context.Context, id int) (*model.Product, error) {
	return p.repository.FindOne(ctx, map[string]any{"id": id})
}

func (p *ProductImpl) UpdateStock(ctx context.Context, id int, stock *requests.ProductStock) (*model.Product, error) {
	tx := p.repository.GetDB(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	product, err := p.repository.FindOne(ctx, map[string]any{"id": id}, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	product.Stock = stock.NewStock
	product, err = p.repository.Update(ctx, product, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return product, nil
}
