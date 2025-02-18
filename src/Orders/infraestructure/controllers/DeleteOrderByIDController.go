package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func (oc *OrderController) DeleteOrder(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := oc.deleteOrderByID.Execute(int32(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}