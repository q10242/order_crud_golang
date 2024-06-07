package controllers

import (
	"net/http"
	"order_crud/database"
	"order_crud/models"

	"github.com/gin-gonic/gin"
)

func FindOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("OrderItem").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func CreateOrder(c *gin.Context) {
	var input struct {
		UserID        uint
		OrderItemID   uint
		Status        string
		Price         float64
		PaymentMethod string
		PaymentStatus string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		UserID:        input.UserID,
		OrderItemID:   input.OrderItemID,
		Status:        input.Status,
		Price:         input.Price,
		PaymentMethod: input.PaymentMethod,
		PaymentStatus: input.PaymentStatus,
	}
	database.DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}
func ShowOrder(c *gin.Context) {
	var order models.Order
	if err := database.DB.Preload("OrderItem").First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func UpdateOrder(c *gin.Context) {
	var input struct {
		UserID        uint
		OrderItemID   uint
		Status        string
		Price         float64
		PaymentMethod string
		PaymentStatus string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	order.UserID = input.UserID
	order.OrderItemID = input.OrderItemID
	order.Status = input.Status
	order.Price = input.Price
	order.PaymentMethod = input.PaymentMethod
	order.PaymentStatus = input.PaymentStatus

	database.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrder(c *gin.Context) {
	var order models.Order
	if err := database.DB.First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": "Order deleted successfully"})
}
