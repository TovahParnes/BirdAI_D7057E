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

	if !response.Data.(models.Err).Success {
		return response
	}

	if response.Data.(*models.User).Active == false {
		return utils.ErrorDeleted("User collection")
	}

	return response
}

func (c *Controller) CCreateUser(user *models.User) (response) {
	coll := c.db.GetCollection(repositories.UserColl)
	createdUser, err := coll.CreateOne(user)
	if createdUser != nil {
		return createdUser.(*models.User), err
	}
	return &models.User{}, err
}

func (c *Controller) CListUsers() ([]*models.User, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.FindAll()
	users := []*models.User{}

	for _, usersObject := range response.Data.([]models.HandlerObject) {
		users = append(users, usersObject.(*models.User))
	}

	return utils.Response(users).Data.([]*models.User)
}

func (c *Controller) CDeleteUser(id string) (*models.User, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	deletedUser, err := coll.DeleteOne(id)
	if deletedUser != nil {
		return deletedUser.(*models.User), err
	}
	return &models.User{}, err
}

func (c *Controller) CUpdateUser(user *models.User) (*models.User, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	updatedUser, err := coll.UpdateOne(bson.D{
		{Key: "_id", Value: user.Id},
		{Key: "username", Value: user.Username},
		{Key: "auth_id", Value: user.AuthId},
		{Key: "created_at", Value: user.CreatedAt},
	})
	if updatedUser != nil {
		return updatedUser.(*models.User), err
	}
	return &models.User{}, err
}
