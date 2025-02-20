package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func (oc *OrderController) DeleteAllOrders(c *gin.Context) {
    if err := oc.deleteAllOrders.Execute(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "All orders deleted successfully"})
}