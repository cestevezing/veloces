package data

import (
	"github.com/cestevezing/veloces/internal/core/model"
	"gorm.io/gorm"
)

type DataLoader struct {
	DB *gorm.DB
}

func NewDataLoader(db *gorm.DB) *DataLoader {
	return &DataLoader{DB: db}
}

func (dl *DataLoader) Load() {
	dl.purgeData()
	dl.loadProducts()
}

func (dl *DataLoader) purgeData() {
	dl.DB.Where("1 = 1").Delete(&model.OrderItem{})
	dl.DB.Where("1 = 1").Delete(&model.Order{})
}
