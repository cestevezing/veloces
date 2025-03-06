package repository

import (
	"errors"

	"github.com/cestevezing/veloces/internal/core/model"
	"github.com/cestevezing/veloces/internal/core/port/repository"
	"gorm.io/gorm"
)

type productImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.IProduct {
	return &productImpl{DB: db}
}

func (p *productImpl) Find() ([]*model.Product, error) {
	var products []*model.Product
	result := p.DB.Find(&products)
	return products, result.Error
}

func (p *productImpl) FindOne(filters map[string]any) (*model.Product, error) {
	product := &model.Product{}
	result := p.DB.Where(filters).First(product)
	if result.Error != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (p *productImpl) Create(product *model.Product) (*model.Product, error) {
	result := p.DB.Create(product)
	if result.Error != nil {
		return nil, errors.New("failed to create product")
	}
	return product, nil
}

func (p *productImpl) Update(product *model.Product) (*model.Product, error) {
	result := p.DB.Save(product)
	if result.Error != nil {
		return nil, errors.New("failed to update product")
	}
	return product, nil
}
