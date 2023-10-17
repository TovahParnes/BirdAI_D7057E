package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"net/http"
)

func (c *Controller) CIsAdmin(authId string) models.Response {
	response := c.CGetUserByAuthId(authId)
	if utils.IsTypeError(response) {
		return response
	}
	/*
	response = c.CGetAdminById(response.Data.(*models.UserDB).Id)
	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.AdminOutput{}) {
		return utils.Response("Is admin")
	}

	return utils.ErrorForbidden("User is not admin")
	*/
	return utils.Response("TODO: Is admin")
}

func (c *Controller) CIsSuperAdmin(authId string) models.Response {
	response := c.CGetUserByAuthId(authId)
	if utils.IsTypeError(response) {
		return response
	}

	/*
	response = c.CGetAdminById(response.Data.(*models.UserDB).Id)
	if utils.IsTypeError(response) {
		return response
	}

	if utils.IsType(response, models.AdminOutput{}) && response.Data.(models.AdminOutput).Access == "superadmin"{
		return utils.Response("Is superadmin")
	}
	
	return utils.ErrorForbidden("User is not superadmin")
	*/
	return utils.Response("TODO: Is superadmin")
}

func (c *Controller) CIsPostsUser(authId string, postId string) models.Response {
	response := c.CGetUserByAuthId(authId)
	if utils.IsTypeError(response) {
		return response
	}
	userId := response.Data.(*models.UserDB).Id

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

	if (response.Data.(*models.UserDB).Id == userId) {
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
	currentUserResponse := c.CIsCurrentUser(authId, userId)
	if currentUserResponse.Data.(models.Err).StatusCode != http.StatusForbidden{
		return currentUserResponse
	}

	adminRresponse := c.CIsAdmin(authId)
	if adminRresponse.Data.(models.Err).StatusCode != http.StatusForbidden{
		return adminRresponse
	}

	if !utils.IsTypeError(currentUserResponse) || !utils.IsTypeError(adminRresponse) {
		return utils.Response("Is current user or admin")
	}

	return utils.ErrorForbidden("User is not current user or admin")
}