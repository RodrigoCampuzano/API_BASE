package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func (uc *UserController) DeleteAllUsers(c *gin.Context) {
    if err := uc.deleteAllUsers.Execute(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})
}