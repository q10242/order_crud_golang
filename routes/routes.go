package routes

import (
	"order_crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.ShowOrder)
	r.PUT("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	r.GET("/orderItems", controllers.FindOrderItems)
	r.POST("/orderItems", controllers.CreateOrderItem)
	r.GET("/orderItems/:id", controllers.ShowOrderItem)
	r.PUT("/orderItems/:id", controllers.UpdateOrderItem)
	r.DELETE("/orderItems/:id", controllers.DeleteOrderItem)

	return r
}
