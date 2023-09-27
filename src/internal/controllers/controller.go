package controllers

import (
	"birdai/src/internal/models"
)

type Controller struct {
	db models.IMongoInstance
}

func NewController(db models.IMongoInstance) Controller {
	return Controller{
		db: db,
	}
}
