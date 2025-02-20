package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)



func (uc *UserController) UpdateUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var userReq UserRequest
    if err := c.ShouldBindJSON(&userReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err = uc.updateUser.Execute(int32(id), userReq.Name, userReq.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}