package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetUserById(id string) models.Response {
	coll := c.db.GetCollection(repositories.UserColl)
	filter := bson.M{"_id": id}
	response := coll.FindOne(filter)

	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.UserOutput{}) && response.Data.(models.UserOutput).Active == false {
		return utils.ErrorDeleted("User collection")
	}

	return response
}

func (c *Controller) CListUsers(set, search string) models.Response {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.FindAll()
	users := []*models.UserOutput{}

	if utils.IsTypeError(response) {
		return response
	}

	for _, usersObject := range response.Data.([]*models.UserOutput) {
		users = append(users, usersObject)
	}

	return utils.Response(users)
}

func (c *Controller) CDeleteUser(id, authId string) models.Response {
	coll := c.db.GetCollection(repositories.UserColl)
	filter := bson.M{"_id": id, "auth_id": authId}
	response := coll.DeleteOne(filter)
	return response
}

func (c *Controller) CUpdateUser(id string, user *models.UserInput) (models.Response) {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.UpdateOne(bson.M{
		"_id": id,
		"username": user.Username,
		"active": user.Active,
	})
	return response
}
