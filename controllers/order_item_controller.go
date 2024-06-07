package controllers

import (
	"net/http"
	"order_crud/database"
	"order_crud/models"

	"github.com/gin-gonic/gin"
)

func FindOrderItems(c *gin.Context) {
	var orderItems []models.OrderItem
	database.DB.Find(&orderItems)
	c.JSON(http.StatusOK, gin.H{"data": orderItems})
}

func CreateOrderItem(c *gin.Context) {
	var input struct {
		ProductName string
		Quantity    int
		Price       float64
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItem := models.OrderItem{
		ProductName: input.ProductName,
		Quantity:    input.Quantity,
		Price:       input.Price,
	}
	database.DB.Create(&orderItem)

	c.JSON(http.StatusOK, gin.H{"data": orderItem})
}

func ShowOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := database.DB.First(&orderItem, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderItem})
}

func UpdateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := database.DB.First(&orderItem, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	var input struct {
		ProductName string
		Quantity    int
		Price       float64
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItem.ProductName = input.ProductName
	orderItem.Quantity = input.Quantity
	orderItem.Price = input.Price

	database.DB.Save(&orderItem)

	c.JSON(http.StatusOK, gin.H{"data": orderItem})
}

func DeleteOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := database.DB.First(&orderItem, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	database.DB.Delete(&orderItem)

	c.JSON(http.StatusOK, gin.H{"data": "Order item deleted"})
}
