package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetUserById(id string) (*models.User, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	user, err := coll.FindOne(id)
	if user != nil {
		return user.(*models.User), err
	}
	return &models.User{}, err
}

func (c *Controller) CCreateUser(user *models.User) (*models.User, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	createdUser, err := coll.CreateOne(user)
	if createdUser != nil {
		return createdUser.(*models.User), err
	}
	return &models.User{}, err
}

func (c *Controller) CListUsers() ([]*models.User, error) {
	coll := c.db.GetCollection(repositories.UserColl)
	usersObjects, err := coll.FindAll()
	users := []*models.User{}
	for _, usersObject := range usersObjects {
		users = append(users, usersObject.(*models.User))
	}

	return users, err
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
