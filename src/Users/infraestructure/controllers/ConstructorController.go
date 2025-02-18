package controllers

import (
    "API/src/Users/application"
)

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserController struct {
	createUser *application.CreateUser
    deleteUserByID *application.DeleteUser
	deleteAllUsers *application.DeleteAllUser
    getAllUsers    *application.GetUsers
    getUserByID    *application.GetUserByID
    updateUser     *application.UpdateUser
}

func NewUserController( createUser *application.CreateUser, updateUser *application.UpdateUser, deleteUserByID *application.DeleteUser,
	deleteAllUsers *application.DeleteAllUser, getAllUsers *application.GetUsers, getUserByID *application.GetUserByID,
	) *UserController { return &UserController{ 
		createUser: createUser, 
		deleteAllUsers: deleteAllUsers, 
		updateUser: updateUser,
        deleteUserByID: deleteUserByID, 
		getAllUsers: getAllUsers, 
		getUserByID: getUserByID,	
	}
}
