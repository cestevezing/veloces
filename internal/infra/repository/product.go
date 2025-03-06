package repository

import (
	"context"
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

func (p *productImpl) Find(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	result := p.DB.Find(&products)
	return products, result.Error
}

func (p *productImpl) FindOne(ctx context.Context, filters map[string]any, tx ...*gorm.DB) (*model.Product, error) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0].WithContext(ctx)
	} else {
		db = p.DB.WithContext(ctx)
	}

	product := &model.Product{}
	result := db.Where(filters).First(product)
	if result.Error != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (p *productImpl) Update(ctx context.Context, product *model.Product, tx ...*gorm.DB) (*model.Product, error) {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0].WithContext(ctx)
	} else {
		db = p.DB.WithContext(ctx)
	}
	result := db.Save(product)
	if result.Error != nil {
		return nil, errors.New("failed to update product")
	}
	return product, nil
}

func (p *productImpl) UpdateStock(ctx context.Context, productID int, quantity int, tx ...*gorm.DB) error {
	var db *gorm.DB
	if len(tx) > 0 {
		db = tx[0].WithContext(ctx)
	} else {
		db = p.DB.WithContext(ctx)
	}

	result := db.Model(&model.Product{}).
		Where("id = ? AND stock >= ?", productID, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity))

	if result.RowsAffected == 0 {
		return errors.New("insufficient stock or concurrent update")
	}
	return nil
}

func (p *productImpl) GetDB(ctx context.Context) *gorm.DB {
	return p.DB
}
