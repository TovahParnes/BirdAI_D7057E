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
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"User ID"
// @Success		200	{object}	models.Response{data=[]models.User}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		410	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router		/users/{id} [get]
func (h *Handler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.controller.CGetUserById(id)

	return utils.ResponseToStatus(c, response)
}

// ListUsers is a function to get a set of all users from database
//
// @Summary		List all users of a specified set
// @Description	List all users of a specified set
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			set	query		int	false	"Set of users"
// @Param			search	query	string	false	"Search parameter for user"
// @Success		200	{object}	models.Response{data=[]models.User}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/list [get]
func (h *Handler) ListUsers(c *fiber.Ctx) error {
	//authId := c.GetReqHeaders()["Authid"]
	//queries := c.Queries()
	//set := queries["set"]
	//search := queries["search"]

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	response := h.controller.CListUsers()
	return utils.ResponseToStatus(c, response)
}

// Login is a function to login a user or create a new user
//
// @Summary		Login a user
// @Description	Login a user or create a new user if there is no existing user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param		set	body		models.User	true	"user"
// @Success		201	{object}	models.Response{data=[]models.Err}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/ [post]
func (h *Handler) LoginUser(c *fiber.Ctx) error {

	//	@Failure	401	{object}	models.Response{}
	//parse auth header
	/*
	auth = h.controller.CAuthenticate(authHeader)
	if !auth.data.success {
		return c.Status(auth.StatusCode).JSON(auth)
	}


	response = h.controller.CLoginUser(auth.data)
	*/


	var user *models.User
	if err := c.BodyParser(&user);
	err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}

	response := h.controller.CLoginUser(user)

	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	return utils.CreationResponseToStatus(c, response)
}

// GetUserMe is a function to get the current user from the databse
//
// @Summary		Get current user
// @Description	Get current user
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	models.Response{data=[]models.User}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		410	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/me [get]
func (h *Handler) GetUserMe(c *fiber.Ctx) error {

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)
	authId := c.GetReqHeaders()["Authid"]

	response := h.controller.CGetUserById(authId)

	return utils.ResponseToStatus(c, response)
}

// UpdateUser is a function to update the given user from the databse
//
// @Summary		Update given user
// @Description	Update given user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param		id	path	string	true	"User ID"
// @Param		user	body		models.User	true	"user"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		403	{object}	models.Response{data=[]models.Err}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/{id} [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	id := c.Params("id")
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		//	@Failure	400	{object}	models.Response{}
		// something with body is wrong/missing
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or user is not the same as the one being updated

	response := h.controller.CUpdateUser(id, user)
	return utils.ResponseToStatus(c, response)
}

// DeleteUser is a function to update the given user from the database
//
// @Summary		Delete given user
// @Description	Delete given user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"User ID"
// @Success		200	{object}	models.Response{}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		403	{object}	models.Response{data=[]models.Err}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	authId := c.GetReqHeaders()["Authid"]

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	response := h.controller.CDeleteUser(id, authId)
	return utils.ResponseToStatus(c, response)
}
