package handlers

import (
	"birdai/src/internal/controllers"
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
)

type Handler struct {
	controller controllers.Controller
	me         *models.UserDB
}

func NewHandler(db repositories.IMongoInstance) Handler {
	return Handler{controller: controllers.NewController(db)}
}
