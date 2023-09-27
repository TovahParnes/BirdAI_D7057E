package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
)

func (c *Controller) CGetUserById(id string) (models.HandlerObject, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	user, err := coll.FindOne(id)
	return user, err
}

func (c *Controller) CCreateUser(user *models.User) error {
	coll := c.db.GetCollection(repositories.UserColl)
	err := coll.CreateOne(user)
	return err
}
