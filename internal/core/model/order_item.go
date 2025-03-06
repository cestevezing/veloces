package model

import "time"

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
