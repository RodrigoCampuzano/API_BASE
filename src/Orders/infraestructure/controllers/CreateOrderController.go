package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func (oc *OrderController) CreateOrder(c *gin.Context) {
    var orderReq OrderRequest
    if err := c.ShouldBindJSON(&orderReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := oc.createOrder.Execute(orderReq.UserID, orderReq.Amount, orderReq.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}