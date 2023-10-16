package controllers

import (
	"birdai/src/internal/repositories"
)

type Controller struct {
	db repositories.IMongoInstance
}

//type IController interface {
//	CGetUserById(authId string) (*models.User, error)
//	CLogin(user *models.User) (*models.User, error)
//	CListUsers(authId string) ([]*models.User, error)
//	CDeleteUser(id, authId string) (*models.User, error)
//	CUpdateUser(user *models.User) (*models.User, error)
//}

func NewController(db repositories.IMongoInstance) Controller {
	return Controller{
		db: db,
	}
}
