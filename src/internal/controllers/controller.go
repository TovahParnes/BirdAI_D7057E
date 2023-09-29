package controllers

import (
	"birdai/src/internal/repositories"
)

type Controller struct {
	db repositories.IMongoInstance
}

func NewController(db repositories.IMongoInstance) Controller {
	return Controller{
		db: db,
	}
}
