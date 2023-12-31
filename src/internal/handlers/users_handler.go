package handlers

// 410 är bra att använda

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// GetUserById is a function to get a user by ID
//
// @Summary		Get user by ID
// @Description	Get user by ID
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"User ID"
// @Success		200	{object}	models.Response{data=models.UserOutput}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		410	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users/{id} [get]
func (h *Handler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	response := utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CGetUserById(id)

	return utils.ResponseToStatus(c, response)
}

// ListUsers is a function to get a set of all users from database
//
// @Summary		List all users of a specified set
// @Description	List all users of a specified set
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		set	query		int	false	"Set of users"
// @Param		search	query	string	false	"Search parameter for user"
// @Success		200	{object}	models.Response{data=[]models.UserOutput}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users/list [get]
func (h *Handler) ListUsers(c *fiber.Ctx) error {
	queries := c.Queries()
	set := queries["set"]
	response := utils.IsValidSet(&set)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	setInt := response.Data.(int)

	search := queries["search"]
	response = utils.IsValidSearch(search)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	response = h.controller.CListUsers(setInt, search)
	return utils.ResponseToStatus(c, response)
}

// LoginUser is a function to login a user or create a new user
//
// @Summary		Login a user
// @Description	Login a user or create a new user if there is no existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		user	body		models.UserLogin	true	"user"
// @Success		200	{object}	models.Response{data=models.UserDB}
// @Success		201	{object}	models.Response{data=models.UserDB}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users [post]
func (h *Handler) LoginUser(c *fiber.Ctx) error {
	var user *models.UserLogin
	if err := c.BodyParser(&user); err != nil {
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response := utils.IsValidUserLogin(user)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.auth.LoginUser(user)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	
	return utils.UserCreationResponseToStatus(c, response)
}

// GetUserMe is a function to get the current user from the databse
//
// @Summary		Get current user
// @Description	Get current user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Success		200	{object}	models.Response{data=models.UserOutput}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		410	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users/me [get]
func (h *Handler) GetUserMe(c *fiber.Ctx) error {

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	userId := response.Data.(models.UserDB).Id

	response = h.controller.CGetUserById(userId)

	return utils.ResponseToStatus(c, response)
}

// UpdateUser is a function to update the given user from the database
//
// @Summary		Update given user
// @Description	Update given user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Param		id	path	string	true	"User ID"
// @Param		user	body		models.UserInput	true	"user"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users/{id} [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	id := c.Params("id")
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CIsCurrentUserOrAdmin(curUserId, id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	var user *models.UserInput
	if err := c.BodyParser(&user); err != nil {
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response = utils.IsValidUserInput(user)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CUpdateUser(id, user)
	return utils.ResponseToStatus(c, response)
}

// DeleteUser is a function to update the given user from the database
//
// @Summary		Delete given user
// @Description	Delete given user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Param		id	path		string	true	"User ID"
// @Success		200	{object}	models.Response{data=string}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	id := c.Params("id")
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CIsCurrentUserOrAdmin(curUserId, id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CDeleteUser(id)
	return utils.ResponseToStatus(c, response)
}
