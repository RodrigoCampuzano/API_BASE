package controllers

import (
   "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func (oc *OrderController) GetOrderByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    order, err := oc.getOrderByID.Execute(int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if order == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}