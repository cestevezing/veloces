package service

import (
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

func (p *ProductImpl) GetAll() ([]*model.Product, error) {
	return p.repository.Find()
}

func (p *ProductImpl) GetByID(id int) (*model.Product, error) {
	return p.repository.FindOne(map[string]any{"id": id})
}

func (p *ProductImpl) UpdateStock(id int, stock *requests.ProductStock) (*model.Product, error) {
	product, err := p.repository.FindOne(map[string]any{"id": id})
	if err != nil {
		return nil, err
	}
	product.Stock = stock.NewStock
	product, err = p.repository.Update(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
