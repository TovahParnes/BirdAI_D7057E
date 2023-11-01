package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetAdminById(id string) (models.Response) {
	response := c.db.Admin.GetAdminById(id)
	if utils.IsTypeError(response) {
		return response
	}
	admin := response.Data.(*models.AdminDB)
	adminResponse := c.cAdminDBToOutput(admin)
	return adminResponse
}

func (c *Controller) CGetAdminByUserId(id string) (models.Response) {
	response := c.db.Admin.GetAdminByUserId(id)
	if utils.IsTypeError(response) {
		return response
	}
	admin := response.Data.(*models.AdminDB)
	adminResponse := c.cAdminDBToOutput(admin)
	return adminResponse
}

func (c *Controller) CListAdmins(set int, search string) (models.Response) {
	response := c.db.Admin.ListAdmins(bson.M{}, set)
	if utils.IsTypeError(response) {
		return response
	}

	output := []*models.AdminOutput{}
	for _, admin := range response.Data.([]models.AdminDB) {
		adminResponse := c.cAdminDBToOutput(&admin)
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
	if utils.IsType(currentAdmin, models.AdminDB{}) {
		return utils.ErrorToResponse(http.StatusConflict, "Admin already exists", "Admin with that user id already exists")
	}

	user := c.db.User.GetUserById(adminInput.UserId)
	if utils.IsTypeError(user) {
		if user.Data.(models.Err).StatusCode == http.StatusNotFound{
			return utils.ErrorNotFoundInDatabase("User with given id does not exist")
		} else {
		return user
		}
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
		response := c.cCheckLastSuperadmin(id)
		if utils.IsTypeError(response) {
			return response
		}
	}
	admin.Id = id
	response := c.db.Admin.UpdateAdmin(*admin)
	return response
}

func (c *Controller) CDeleteAdmin(id string) (models.Response) {
	response := c.cCheckLastSuperadmin(id)
	if utils.IsTypeError(response) {
		return response
	}
	response = c.db.Admin.DeleteAdmin(id)
	return response
}

func (c *Controller) cCheckLastSuperadmin(id string) (models.Response) {
	filter := bson.M{"access": "superadmin"}
	response := c.db.Admin.ListAdmins(filter, 0)
	if utils.IsTypeError(response) {
		return response
	}
	admins := response.Data.([]models.AdminDB)

	if len(admins) == 1 && admins[0].Id == id{
		return utils.ErrorToResponse(http.StatusConflict, "Cannot remove last superadmin", "There must be at least one superadmin")
	}

	return utils.Response(nil)
}

func (c *Controller) cAdminDBToOutput(admin *models.AdminDB) (models.Response) {
	userResponse := c.db.User.GetUserById(admin.UserId)
	if utils.IsTypeError(userResponse) {
		return userResponse
	}
	userDB := userResponse.Data.(*models.UserDB)
	userOutput := models.UserDBToOutput(userDB)
	adminOutput := models.AdminDBToOutput(admin, userOutput)
	return utils.Response(adminOutput)
}