package service

import (
	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/model"
)

type IProduct interface {
	GetAll() ([]*model.Product, error)
	GetByID(id int) (*model.Product, error)
	UpdateStock(id int, stock *requests.ProductStock) (*model.Product, error)
}
