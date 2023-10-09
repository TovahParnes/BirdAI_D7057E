package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetUserById(id string) (models.Response) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.FindOne(id)

	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsTypeUser(response) && response.Data.(models.User).Active == false {
		return utils.ErrorDeleted("User collection")
	}

	return response
}

func (c *Controller) CCreateUser(user *models.User) (models.Response) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.CreateOne(user)
	if utils.IsTypeError(response) {
		return response
	}
	return response
}

func (c *Controller) CListUsers() (models.Response) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.FindAll()
	users := []*models.User{}

	if utils.IsTypeError(response) {
		return response
	}

	for _, usersObject := range response.Data.([]models.HandlerObject) {
		users = append(users, usersObject.(*models.User))
	}

	return utils.Response(users)
}

func (c *Controller) CDeleteUser(id string) (models.Response) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.DeleteOne(id)
	return response
}

func (c *Controller) CUpdateUser(id string, user *models.User) (models.Response) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.UpdateOne(bson.D{
		{Key: "_id", Value: id},
		{Key: "username", Value: user.Username},
		{Key: "auth_id", Value: user.AuthId},
		{Key: "created_at", Value: user.CreatedAt},
		{Key: "active", Value: user.Active},
	})
	return response
}
