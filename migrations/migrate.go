package migrations

import (
	"order_crud/database"
	"order_crud/models"
)

func Migrate() {
	database.DB.AutoMigrate(&models.Order{}, &models.OrderItem{})
}
