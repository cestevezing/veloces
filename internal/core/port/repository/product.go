package repository

import (
	"context"

	"github.com/cestevezing/veloces/internal/core/model"
	"gorm.io/gorm"
)

type IProduct interface {
	Find(ctx context.Context) ([]*model.Product, error)
	FindOne(ctx context.Context, filters map[string]any, tx ...*gorm.DB) (*model.Product, error)
	Update(ctx context.Context, product *model.Product, tx ...*gorm.DB) (*model.Product, error)
	UpdateStock(ctx context.Context, productID int, quantity int, tx ...*gorm.DB) error
	GetDB(ctx context.Context) *gorm.DB
}
