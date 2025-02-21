package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (oc *OrderController) UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var orderReq OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = oc.updateOrder.Execute(int32(id), orderReq.UserID, orderReq.Amount, orderReq.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la orden"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orden actualizada exitosamente"})
}
