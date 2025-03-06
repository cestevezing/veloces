package repository

import (
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

func (o *OrderImpl) Create(order *model.Order) (*model.Order, error) {
	result := o.DB.Create(order)
	if result.Error != nil {
		return nil, errors.New("error creating order")
	}
	return order, nil
}

func (o *OrderImpl) FindOne(filter map[string]any) (*model.Order, error) {
	order := &model.Order{}
	result := o.DB.Preload("Items").Where(filter).First(order)
	if result.Error != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (o *OrderImpl) Update(order *model.Order) (*model.Order, error) {
	result := o.DB.Save(order)
	if result.Error != nil {
		return nil, errors.New("error updating order")
	}
	return order, nil
}
