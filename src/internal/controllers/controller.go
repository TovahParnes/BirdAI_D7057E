package controllers

import (
	"birdai/src/internal/repositories"
)

type Controller struct {
	db repositories.RepositoryEndpoints
}

func NewController(db repositories.RepositoryEndpoints) Controller {
	return Controller{
		db: db,
	}
}
