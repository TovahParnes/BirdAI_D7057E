package handlers

import (
	"birdai/src/internal/authentication"
	"birdai/src/internal/controllers"
	"birdai/src/internal/repositories"
)

type Handler struct {
	controller controllers.Controller
	auth       authentication.Authentication
}

func NewHandler(db repositories.RepositoryEndpoints) Handler {
	return Handler{
		controller: controllers.NewController(db),
		auth:       authentication.NewAuthentication(db.User),
	}
}
