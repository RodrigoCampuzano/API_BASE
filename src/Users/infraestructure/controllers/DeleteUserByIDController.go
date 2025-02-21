package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }
    user, err := uc.getUserByID.Execute(int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }
    ordersExist, err := uc.ordersExistForUser(int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar las órdenes del usuario"})
        return
    }
    if ordersExist {
        c.JSON(http.StatusConflict, gin.H{"error": "No se puede eliminar el usuario porque existen órdenes asociadas"})
        return
    }
    if err := uc.deleteUserByID.Execute(int32(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
