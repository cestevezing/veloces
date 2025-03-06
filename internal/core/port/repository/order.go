package repository

import (
	"context"

	"github.com/cestevezing/veloces/internal/core/model"
	"gorm.io/gorm"
)

type IOrder interface {
	Create(ctx context.Context, order *model.Order, tx ...*gorm.DB) (*model.Order, error)
	FindOne(ctx context.Context, filter map[string]any) (*model.Order, error)
	GetDB(ctx context.Context) *gorm.DB
}
