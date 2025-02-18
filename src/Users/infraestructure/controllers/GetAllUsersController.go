package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)



func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.getAllUsers.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}