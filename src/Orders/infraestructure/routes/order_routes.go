package routes

import (
	"API/src/Orders/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderController *controllers.OrderController) {
    router.POST("/orders",orderController.CreateOrder) 
    router.GET("/orders", orderController.GetAllOrders)
    router.GET("/orders/:id", orderController.GetOrderByID)
    router.PUT("/orders/:id", orderController.UpdateOrder)
    router.DELETE("/orders/:id", orderController.DeleteOrder)
    router.DELETE("/orders", orderController.DeleteAllOrders)
}