package controllers

import (
   "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func (oc *OrderController) GetOrderByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    order, err := oc.getOrderByID.Execute(int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la orden"})
        return
    } 
    if order == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
        return
    }
    c.JSON(http.StatusOK, order)
}