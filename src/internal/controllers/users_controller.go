package controllers

import (
	"birdai/src/internal/storage"
)

func (c *Controller) CGetUserById(id string) (storage.HandlerObject, error) {
	coll := c.db.GetCollection(storage.UserColl)
	user, err := coll.FindOne(id)
	return user, err
}

func (c *Controller) CCreateUser(user *storage.User) error {
	coll := c.db.GetCollection(storage.UserColl)
	err := coll.CreateOne(user)
	return err
}
