package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductName string
	Quantity    int
	Price       float64
}
