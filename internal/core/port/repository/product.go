package repository

import "github.com/cestevezing/veloces/internal/core/model"

type IProduct interface {
	Find() ([]*model.Product, error)
	FindOne(filters map[string]any) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
}
