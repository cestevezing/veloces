package data

import (
	"github.com/cestevezing/veloces/internal/core/model"
)

func (dl *DataLoader) loadProducts() error {
	defaultProducts := []*model.Product{
		{Name: "Laptop", Price: 999.99, Stock: 10},
		{Name: "Smartphone", Price: 499.99, Stock: 25},
		{Name: "Tablet", Price: 299.99, Stock: 15},
		{Name: "Headphones", Price: 99.99, Stock: 50},
		{Name: "Keyboard", Price: 49.99, Stock: 30},
	}

	for _, v := range defaultProducts {
		if err := dl.DB.Where("name =?", v.Name).First(&v).Error; err != nil {
			if err := dl.DB.Create(&v).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
