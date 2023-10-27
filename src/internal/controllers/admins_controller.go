package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetAdminById(id string) (models.Response) {
	response := c.db.Admin.GetAdminById(id)
	if utils.IsTypeError(response) {
		return response
	}
	admin := response.Data.(*models.AdminDB)
	adminResponse := c.AdminDBToOutput(admin)
	return adminResponse
}

func (c *Controller) CGetAdminByUserId(id string) (models.Response) {
	response := c.db.Admin.GetAdminByUserId(id)
	if utils.IsTypeError(response) {
		return response
	}
	admin := response.Data.(*models.AdminDB)
	adminResponse := c.AdminDBToOutput(admin)
	return adminResponse
}

func (c *Controller) CListAdmins(set int, search string) (models.Response) {
	response := c.db.Admin.ListAdmins(bson.M{}, set)
	if utils.IsTypeError(response) {
		return response
	}

	output := []*models.AdminOutput{}
	for _, admin := range response.Data.([]models.AdminDB) {
		adminResponse := c.AdminDBToOutput(&admin)
		if utils.IsTypeError(adminResponse) {
			return adminResponse
		}
		
		output = append(output, adminResponse.Data.(*models.AdminOutput))
	}

	return utils.Response(output)
}

func (c *Controller) CCreateAdmin(adminInput *models.AdminInput) (models.Response) {
	currentAdmin := c.db.Admin.GetAdminByUserId(adminInput.UserId)
	if utils.IsTypeError(currentAdmin) && currentAdmin.Data.(models.Err).StatusCode != http.StatusNotFound{
		return currentAdmin
	}
	if utils.IsType(currentAdmin, models.AdminOutput{}) {
		return utils.ErrorToResponse(http.StatusConflict, "Admin already exists", "Admin with that user id already exists")
	}

	admin := &models.AdminDB{
		UserId: adminInput.UserId,
		Access: adminInput.Access,
	}
	response := c.db.Admin.CreateAdmin(*admin)
	return response
}

func (c *Controller) CUpdateAdmin(id string, admin *models.AdminInput) (models.Response) {
	if (admin.Access == "admin") {
		response := c.CCheckLastSuperadmin(id)
		if utils.IsTypeError(response) {
			return response
		}
	}
	admin.Id = id
	response := c.db.Admin.UpdateAdmin(*admin)
	return response
}

func (c *Controller) CDeleteAdmin(id string) (models.Response) {
	response := c.CCheckLastSuperadmin(id)
	if utils.IsTypeError(response) {
		return response
	}
	response = c.db.Admin.DeleteAdmin(id)
	return response
}

func (c *Controller) CCheckLastSuperadmin(id string) (models.Response) {
	filter := bson.M{"access": "superadmin"}
	response := c.db.Admin.ListAdmins(filter, 0)
	if utils.IsTypeError(response) {
		return response
	}
	admins := response.Data.([]models.AdminDB)
	fmt.Println("superadmin list: ",admins)
	fmt.Println("id: ", id)

	if len(admins) == 1 && admins[0].Id == id{
		fmt.Println("test in if")
		return utils.ErrorToResponse(http.StatusConflict, "Cannot remove last superadmin", "There must be at least one superadmin")
	}

	return utils.Response(nil)
}

func (c *Controller) AdminDBToOutput(admin *models.AdminDB) (models.Response) {
	userResponse := c.db.User.GetUserById(admin.UserId)
	if utils.IsTypeError(userResponse) {
		return userResponse
	}
	userDB := userResponse.Data.(*models.UserDB)
	userOutput := models.UserDBToOutput(*userDB)
	adminOutput := models.AdminDBToOutput(admin, userOutput)
	return utils.Response(adminOutput)
}

func (c *Controller) FirstAdmin(curUserId string) models.Response {
	response := c.db.Admin.ListAdmins(bson.M{}, 0)
	if utils.IsTypeError(response) {
		return response
	}
	if len(response.Data.([]models.AdminDB)) == 0 {
		response = c.CCreateAdmin(&models.AdminInput{
			UserId: curUserId,
			Access: "superadmin",
		},)
		fmt.Println("create admin response: ",response)
		return response
	}
	return response
}

func (c *Controller) ResetAdmin() models.Response {
	currentAdmins := c.db.Admin.ListAllAdmins(0)
	if utils.IsTypeError(currentAdmins) && currentAdmins.Data.(models.Err).StatusCode != http.StatusNotFound{
		return currentAdmins
	}
	if len(currentAdmins.Data.([]models.AdminDB)) != 0 {
		for _, admin := range currentAdmins.Data.([]models.AdminDB) {
			c.db.Admin.DeleteAdmin(admin.Id)
		}
	}
	return utils.Response(nil)
}