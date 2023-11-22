package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetUserById(id string) models.Response {
	response := c.db.User.GetUserById(id)

	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.UserDB{}) && response.Data.(*models.UserDB).Active == false {
		return utils.ErrorDeleted("User collection")
	}

	db := response.Data.(*models.UserDB)
	output := models.UserDBToOutput(db)
	return utils.Response(output)
}

func (c *Controller) CGetUserByAuthId(authId string) models.Response {
	response := c.db.User.GetUserByAuthId(authId)

	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.UserDB{}) && response.Data.(*models.UserDB).Active == false {
		return utils.ErrorDeleted("User collection")
	}

	db := response.Data.(*models.UserDB)
	output := models.UserDBToOutput(db)
	return utils.Response(output)
}

func (c *Controller) CListUsers(set int) models.Response {
	response := c.db.User.ListUsers(bson.M{}, set)
	users := []*models.UserOutput{}

	if utils.IsTypeError(response) {
		return response
	}

	for _, usersObject := range response.Data.([]models.UserDB) {
		users = append(users, models.UserDBToOutput(&usersObject))
	}

	return utils.Response(users)
}

func (c *Controller) CDeleteUser(id string) models.Response {
	return c.db.User.DeleteUser(id)
}

func (c *Controller) CUpdateUser(id string, user *models.UserInput) models.Response {
	response := c.db.User.UpdateUser(id, *user)
	if utils.IsTypeError(response) {
		return response
	}

	db := response.Data.(*models.UserDB)
	output := models.UserDBToOutput(db)
	return utils.Response(output)
}
