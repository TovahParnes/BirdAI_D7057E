package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"net/http"
)

func (c *Controller) CIsAdmin(curUserId string) models.Response {
	response := c.CGetAdminByUserId(curUserId)
	fmt.Println("Admin response:",response.Data)
	if utils.IsTypeError(response) && response.Data.(models.Err).StatusCode != http.StatusNotFound{
		fmt.Println("if 1")
		return response
	}

	if utils.IsType(response, models.AdminOutput{}){
		fmt.Println("if 2")
		return utils.Response("Is admin")
	}
	
	return utils.ErrorForbidden("User is not admin")
}

func (c *Controller) CIsSuperAdmin(curUserId string) models.Response {
	response := c.CGetAdminByUserId(curUserId)
	if utils.IsTypeError(response) && response.Data.(models.Err).StatusCode != http.StatusNotFound{
		return response
	}

	if utils.IsType(response, models.AdminOutput{}) && response.Data.(*models.AdminOutput).Access == "superadmin"{
		return utils.Response("Is superadmin")
	}
	
	return utils.ErrorForbidden("User is not superadmin")
}

func (c *Controller) CIsPostsUser(curUserId string, postId string) models.Response {
	response := c.CGetPostById(postId)
	if utils.IsTypeError(response)  && response.Data.(models.Err).StatusCode != http.StatusNotFound{
		return response
	}

	if utils.IsType(response, models.PostOutput{}) && response.Data.(models.PostOutput).User.Id == curUserId {
		return utils.Response("Is posts user")
	}

	return utils.ErrorForbidden("User is not posts user")
}

func (c *Controller) CIsCurrentUser(curUserId string, userId string) models.Response {
	if (curUserId == userId) {
		return utils.Response("Is current user")
	}

	return utils.ErrorForbidden("User is not current user")
}

func (c *Controller) CIsPostsUserOrAdmin(curUserId string, postId string) models.Response {
	postResponse := c.CIsPostsUser(curUserId, postId)
	if utils.IsTypeError(postResponse) && postResponse.Data.(models.Err).StatusCode != http.StatusForbidden {
		return postResponse
	}
	if !utils.IsTypeError(postResponse) {
		return utils.Response("Is posts user")
	}

	adminRresponse := c.CIsAdmin(curUserId)
	if utils.IsTypeError(adminRresponse) && adminRresponse.Data.(models.Err).StatusCode != http.StatusForbidden  {
		return adminRresponse
	}
	if !utils.IsTypeError(adminRresponse) {
		return utils.Response("Is admin")
	}

	return utils.ErrorForbidden("User is not posts user or admin")
}

func (c *Controller) CIsCurrentUserOrAdmin(curUserId string, userId string) models.Response {
	currentUserResponse := c.CIsCurrentUser(curUserId, userId)
	//fmt.Println("Curtent user response:",currentUserResponse.Data)
	//fmt.Println("Curtent user response:",currentUserResponse.Data.(models.Err).StatusCode)
	//fmt.Println(utils.IsTypeError(currentUserResponse) && currentUserResponse.Data.(models.Err).StatusCode != http.StatusForbidden)

	if utils.IsTypeError(currentUserResponse) && currentUserResponse.Data.(models.Err).StatusCode != http.StatusForbidden{
		//fmt.Println("if 1")
		return currentUserResponse
	}
	if !utils.IsTypeError(currentUserResponse) {
		//fmt.Println("if 2")
		return utils.Response("Is current user")
	}

	adminRresponse := c.CIsAdmin(curUserId)
	fmt.Println("Admin response:",adminRresponse.Data)
	if utils.IsTypeError(adminRresponse) && adminRresponse.Data.(models.Err).StatusCode != http.StatusForbidden{
		fmt.Println("if 3")
		return adminRresponse
	}
	if !utils.IsTypeError(adminRresponse) {
		fmt.Println("if 4")
		return utils.Response("Is admin")
	}

	return utils.ErrorForbidden("User is not current user or admin")
}

func (c *Controller) CIsCurrentUserOrSuperAdmin(curUserId string, userId string) models.Response {
	currentUserResponse := c.CIsCurrentUser(curUserId, userId)
	if utils.IsTypeError(currentUserResponse) && currentUserResponse.Data.(models.Err).StatusCode != http.StatusForbidden{
		return currentUserResponse
	}
	if !utils.IsTypeError(currentUserResponse) {
		return utils.Response("Is current user")
	}

	superAdminRresponse := c.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(superAdminRresponse) && superAdminRresponse.Data.(models.Err).StatusCode != http.StatusForbidden{
		return superAdminRresponse
	}
	if !utils.IsTypeError(superAdminRresponse) {
		return utils.Response("Is superadmin")
	}

	return utils.ErrorForbidden("User is not current user or superadmin")
}