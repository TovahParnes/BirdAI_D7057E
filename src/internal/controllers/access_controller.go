package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
)

func (c *Controller) CIsAdmin(authId string) models.Response {
	response := c.CGetUserById(authId)
	if utils.IsTypeError(response) {
		return response
	}

	response = c.CGetAdminById(response.Data.(*models.UserOutput).Id)
	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.AdminOutput{}) {
		return utils.Response("Is admin")
	}

	return utils.ErrorForbidden("User is not admin")
}

func CGetUserByAuthId(authId string) {
	panic("unimplemented")
}

func (c *Controller) CIsSuperAdmin(authId string) models.Response {
	response := c.CGetUserByAuthId(authId)
	if utils.IsTypeError(response) {
		return response
	}

	response = c.CGetAdminById(response.Data.(*models.UserOutput).Id)
	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.AdminOutput{}) && response.Data.(models.AdminOutput).Access == "superadmin"{
		return utils.Response("Is superadmin")
	}
	
	return utils.ErrorForbidden("User is not superadmin")
}

func (c *Controller) CIsPostsUser(authId string, postId string) models.Response {
	response := c.CGetUserByAuthId(authId)
	if utils.IsTypeError(response) {
		return response
	}
	userId := response.Data.(*models.UserOutput).Id

	response = c.CGetPostById(postId)
	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.PostOutput{}) && response.Data.(models.PostOutput).User.Id == userId {
		return utils.Response("Is posts user")
	}

	return utils.ErrorForbidden("User is not posts user")
}

func (c *Controller) CIsCurrentUser(authId string, userId string) models.Response {
	response := c.CGetUserByAuthId(authId)
	if utils.IsTypeError(response) {
		return response
	}

	if (response.Data.(*models.UserOutput).Id == userId) {
		return utils.Response("Is current user")
	}

	return utils.ErrorForbidden("User is not current user")
}

func (c *Controller) CIsPostsUserOrAdmin(authId string, userId string) models.Response {
	postResponse := c.CIsPostsUser(authId, userId)
	if utils.IsTypeError(postResponse) {
		return postResponse
	}

	adminRresponse := c.CIsAdmin(authId)
	if utils.IsTypeError(adminRresponse) {
		return adminRresponse
	}

	if utils.IsType(postResponse, models.PostOutput{}) || utils.IsType(adminRresponse, models.AdminOutput{}) {
		return utils.Response("Is posts user or admin")
	}

	return utils.ErrorForbidden("User is not current user")
}

func (c *Controller) CIsCurrentUserOrAdmin(authId string, userId string) models.Response {
	CurrentUserResponse := c.CIsCurrentUser(authId, userId)
	if utils.IsTypeError(CurrentUserResponse) {
		return CurrentUserResponse
	}

	adminRresponse := c.CIsAdmin(authId)
	if utils.IsTypeError(adminRresponse) {
		return adminRresponse
	}

	if utils.IsType(CurrentUserResponse, models.UserOutput{}) || utils.IsType(adminRresponse, models.AdminOutput{}) {
		return utils.Response("Is current user or admin")
	}

	return utils.ErrorForbidden("User is not current user or admin")
}