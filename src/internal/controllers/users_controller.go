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

func (c *Controller) CListUsers() models.Response {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.FindAll(bson.M{}, 0, 0)
	users := []*models.UserOutput{}

	if utils.IsTypeError(response) {
		return response
	}

	for _, usersObject := range response.Data.([]*models.UserDB) {
		users = append(users, UserDBToOutput(usersObject))
	}

	return utils.Response(users)
}

func (c *Controller) CDeleteUser(id, authId string) models.Response {
	coll := c.db.GetCollection(repositories.UserColl)
	filter := bson.M{"_id": id, "auth_id": authId}
	response := coll.DeleteOne(filter)
	return response
}

func (c *Controller) CUpdateUser(id string, user *models.UserInput) models.Response {
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.UpdateOne(bson.M{
		"_id":      id,
		"username": user.Username,
		"active":   user.Active,
	})
	return response
}

func UserDBToOutput(db *models.UserDB) *models.UserOutput {
	return &models.UserOutput{
		Id:        db.Id,
		Username:  db.Username,
		CreatedAt: db.CreatedAt,
		Active:    db.Active,
	}
}
