package handlers

import (
	"birdai/src/internal/controllers"
	"birdai/src/internal/storage"
)

type Handler struct {
	controller controllers.Controller
}

func NewHandler(db storage.IMongoInstance) Handler {
	return Handler{controller: controllers.NewController(db)}
}
