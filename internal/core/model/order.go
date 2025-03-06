package model

import "time"

type Order struct {
	ID           uint    `gorm:"primaryKey"`
	CustomerName string  `gorm:"not null"`
	TotalAmount  float64 `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Items        []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}
