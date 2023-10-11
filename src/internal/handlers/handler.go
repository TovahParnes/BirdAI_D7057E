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

func NewHandler(db repositories.IMongoInstance) Handler {
	return Handler{
		controller: controllers.NewController(db),
		auth:       authentication.NewAuthentication(db.GetCollection(repositories.UserColl)),
	}
}
