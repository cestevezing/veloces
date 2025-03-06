package service

import (
	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/model"
)

type IOrder interface {
	Create(order *requests.OrderCreate) (*model.Order, error)
	GetByID(id int) (*model.Order, error)
}
