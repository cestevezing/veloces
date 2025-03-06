package repository

import (
	"context"
	"errors"

	"github.com/cestevezing/veloces/internal/core/model"
	"github.com/cestevezing/veloces/internal/core/port/repository"
	"gorm.io/gorm"
)

type OrderImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.IOrder {
	return &OrderImpl{DB: db}
}

func (o *OrderImpl) Create(ctx context.Context, order *model.Order, tx ...*gorm.DB) (*model.Order, error) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0].WithContext(ctx)
	} else {
		db = o.DB.WithContext(ctx)
	}
	result := db.Create(order)
	if result.Error != nil {
		return nil, errors.New("error creating order")
	}
	return order, nil
}

func (o *OrderImpl) FindOne(ctx context.Context, filter map[string]any) (*model.Order, error) {
	order := &model.Order{}
	result := o.DB.Preload("Items").Where(filter).First(order)
	if result.Error != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (o *OrderImpl) GetDB(ctx context.Context) *gorm.DB {
	return o.DB
}
