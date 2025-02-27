package application

import (
	"API/src/Users/domain/entities"
	"API/src/Users/domain/repositories"
)

type GetUsers struct {
	db repositories.IntrUser
}

func NewGetUser(db repositories.IntrUser) *GetUsers {
	return &GetUsers{db: db}
}

func (vu *GetUsers) Execute() ([]*entities.User, error){
	return vu.db.GetAll()
}
