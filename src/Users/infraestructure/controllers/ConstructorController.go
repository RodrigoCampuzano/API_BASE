package controllers

import (
	"API/src/Users/application"
)

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserController struct {
	createUser         *application.CreateUser
	deleteUserByID     *application.DeleteUser
	getAllUsers        *application.GetUsers
	getUserByID        *application.GetUserByID
	updateUser         *application.UpdateUser
	ordersExistForUser func(userID int32) (bool, error)
}

func NewUserController(createUser *application.CreateUser, updateUser *application.UpdateUser, deleteUserByID *application.DeleteUser,
	getAllUsers *application.GetUsers, getUserByID *application.GetUserByID, ordersExistForUser func(userID int32) (bool, error),
) *UserController {
	return &UserController{
		createUser:         createUser,
		updateUser:         updateUser,
		deleteUserByID:     deleteUserByID,
		getAllUsers:        getAllUsers,
		getUserByID:        getUserByID,
		ordersExistForUser: ordersExistForUser,
	}
}
