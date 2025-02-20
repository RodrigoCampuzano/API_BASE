package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func (oc *OrderController) GetAllOrders(c *gin.Context) {
    orders, err := oc.getAllOrders.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}