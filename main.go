package main

import (
    orderApp "API/src/Orders/application"
    oderCtrl "API/src/Orders/infraestructure/controllers"
    infraRepoOrder "API/src/Orders/infraestructure/repositories"
    orderRoutes "API/src/Orders/infraestructure/routes"

    userApp "API/src/Users/application"
    userCtrl "API/src/Users/infraestructure/controllers"
    infraRepoUser "API/src/Users/infraestructure/repositories"
    userRoutes "API/src/Users/infraestructure/routes"
            
    "API/src/core/middleware"
    "API/src/core/db"
    "github.com/gin-gonic/gin"
)

func main() {
    database := db.ConnectionDB()
    defer database.Close()

    userRepo := infraRepoUser.NewUserRepository(database)
    orderRepo := infraRepoOrder.NewOrderRepository(database)

    createUser := userApp.NewCreateUser(userRepo)
    getAllUsers := userApp.NewGetUser(userRepo)
    getUserByID := userApp.NewGetUserByID(userRepo)
    updateUser := userApp.NewUpdateUser(userRepo)
    deleteUserByID := userApp.NewDeleteUserByID(userRepo)
    deleteAllUsers := userApp.NewDeleteAllUser(userRepo)

    createOrder := orderApp.NewCreateOrder(orderRepo)
    getAllOrders := orderApp.NewGetAllOrders(orderRepo)
    getOrderByID := orderApp.NewGetOrderByID(orderRepo)
    updateOrder := orderApp.NewUpdateOrder(orderRepo)
    deleteOrderByID := orderApp.NewDeleteOrderByID(orderRepo)
    deleteAllOrders := orderApp.NewDeleteAllOrders(orderRepo)

    userController := userCtrl.NewUserController(createUser, updateUser, deleteUserByID, deleteAllUsers, getAllUsers, getUserByID)
    orderController := oderCtrl.NewOrderController(createOrder, updateOrder, getOrderByID, getAllOrders, deleteAllOrders, deleteOrderByID)

    router := gin.Default()

    router.Use(middleware.CORSConfig())

    userRoutes.SetupRoutes(router, userController)
    orderRoutes.SetupOrderRoutes(router, orderController    )

    router.Run("127.0.0.1:8080")
}
