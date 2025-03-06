package service

import (
	"context"

	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/model"
)

type IProduct interface {
	GetAll(ctx context.Context) ([]*model.Product, error)
	GetByID(ctx context.Context, id int) (*model.Product, error)
	UpdateStock(ctx context.Context, id int, stock *requests.ProductStock) (*model.Product, error)
}
