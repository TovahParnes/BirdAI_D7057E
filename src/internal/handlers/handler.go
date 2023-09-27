package handlers

import (
	"birdai/src/internal/controllers"
	"birdai/src/internal/models"
)

type Handler struct {
	controller controllers.Controller
}

func NewHandler(db models.IMongoInstance) Handler {
	return Handler{controller: controllers.NewController(db)}
}
