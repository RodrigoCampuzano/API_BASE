package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) CreateUser(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.createUser.Execute(userRequest.Name, userRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}
