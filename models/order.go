package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        uint
	OrderItemID   uint
	Status        string    `gorm:"default:'pending'"`
	Price         float64   `gorm:"default:0"`
	PaymentMethod string    `gorm:"default:''"`
	PaymentStatus string    `gorm:"default:''"`
	OrderItem     OrderItem `gorm:"foreignKey:OrderItemID"`
}
