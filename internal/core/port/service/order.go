package service

import (
	"context"

	"github.com/cestevezing/veloces/internal/core/dto/requests"
	"github.com/cestevezing/veloces/internal/core/dto/response"
	"github.com/cestevezing/veloces/internal/core/model"
)

type IOrder interface {
	Create(ctx context.Context, order *requests.OrderCreate) (*model.Order, error)
	GetByID(ctx context.Context, id int) (*response.Order, error)
}
