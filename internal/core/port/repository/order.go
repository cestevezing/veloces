package repository

import (
	"github.com/cestevezing/veloces/internal/core/model"
)

type IOrder interface {
	Create(order *model.Order) (*model.Order, error)
	FindOne(filter map[string]any) (*model.Order, error)
	Update(order *model.Order) (*model.Order, error)
}
