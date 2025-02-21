package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (oc *OrderController) GetAllOrders(c *gin.Context) {
	orders, err := oc.getAllOrders.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las órdenes"})
		return
	}
	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No existen órdenes"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
