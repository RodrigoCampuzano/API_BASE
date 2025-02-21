package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.getAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No existen usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}
