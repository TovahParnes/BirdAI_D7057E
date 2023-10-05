package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetUserById(authId string) (*models.User, error) {
	user, err := c.auth.CheckUser(authId)
	if err != nil {
		return &models.User{}, err
	}
	return user.(*models.User), err
}

func (c *Controller) CLogin(user *models.User) (string, error) {
	authUser, err := c.auth.LoginUser(user)
	if err != nil {
		return "", err
	}
	return authUser, nil
}

func (c *Controller) CListUsers(authId string) ([]*models.User, error) {
	_, err := c.auth.CheckUser(authId)
	if err != nil {
		return nil, err
	}
	coll := c.db.GetCollection(repositories.UserColl)
	inter, err := coll.FindAll()
	users := inter.([]*models.User)

	return users, err
}

func (c *Controller) CDeleteUser(id, authId string) (*models.User, error) {
	_, err := c.auth.CheckUser(authId)
	if err != nil {
		return nil, err
	}
	coll := c.db.GetCollection(repositories.UserColl)
	deletedUser, err := coll.DeleteOne(bson.M{"_id": id})
	if deletedUser != nil {
		return deletedUser.(*models.User), err
	}
	return &models.User{}, err
}

func (c *Controller) CUpdateUser(user *models.User) (*models.User, error) {
	_, err := c.auth.CheckUser(user.AuthId)
	if err != nil {
		return nil, err
	}
	coll := c.db.GetCollection(repositories.UserColl)
	updatedUser, err := coll.UpdateOne(bson.M{
		"_id":        user.Id,
		"username":   user.Username,
		"auth_id":    user.AuthId,
		"created_at": user.CreatedAt,
	})
	if updatedUser != nil {
		return updatedUser.(*models.User), err
	}
	return &models.User{}, err
}
