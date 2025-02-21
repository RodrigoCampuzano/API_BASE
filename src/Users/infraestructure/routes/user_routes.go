package routes

import (
    "API/src/Users/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
    router.POST("/users", userController.CreateUser)
    router.GET("/users", userController.GetAllUsers)
    router.GET("/users/:id", userController.GetUserByID)
    router.PUT("/users/:id", userController.UpdateUser)
    router.DELETE("/users/:id", userController.DeleteUser)
}