package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"net/http"
)

func (c *Controller) CGetAdminById(id string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.AdminColl)
	filter := bson.M{"_id": id}
	response := coll.FindOne(filter)
	return response
	*/
	return utils.ErrorNotImplemented("CGetAdminById")
}

func (c *Controller) CListAdmins(set string, search string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.AdminColl)
	response := coll.FindAll()
	return response
	*/
	return utils.ErrorNotImplemented("CListAdmins")
}

func (c *Controller) CCreateAdmin(authId string, admin *models.AdminInput) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.AdminColl)
	response := coll.CreateOne(admin)
	return response
	*/
	return utils.ErrorNotImplemented("CCreateAdmin")
}

func (c *Controller) CUpdateAdmin(userId string, admin *models.AdminInput) (models.Response) {
	//check if updating last superadmin to admin

	/*
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.UpdateOne(bson.M{
		"_id": admin.Id,
		"user_id": userId,
		"access": admin.access,
	})
	return response
	*/

	
	return utils.ErrorNotImplemented("CUpdateAdmin")
}

func (c *Controller) CDeleteAdmin(id string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.AdminColl)
	filter := bson.M{"_id": id}
	response := coll.DeleteOne(filter)
	return response
	*/
	return utils.ErrorNotImplemented("CDeleteAdmin")
}

func (c *Controller) CCheckLastSuperadmin() (models.Response) {
	response := c.CListAdmins("1", "Access: superadmin")
	if utils.IsTypeError(response) {
		return response
	}

	admins := response.Data.([]models.AdminOutput)
	if len(admins) <= 1{
		return utils.ErrorToResponse(http.StatusConflict, "Cannot delete last superadmin", "There must be at least one superadmin")
	}

	return utils.ErrorNotImplemented("CDeleteAdmin")
}