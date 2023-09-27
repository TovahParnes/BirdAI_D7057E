package controllers

import (
	"birdai/src/internal/storage"
)

type Controller struct {
	db storage.IMongoInstance
}

func NewController(db storage.IMongoInstance) Controller {
	return Controller{
		db: db,
	}
}
