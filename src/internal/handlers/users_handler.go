package handlers

// 410 är bra att använda

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetUserById is a function to get a user by ID
//
// @Summary		Get user by ID
// @Description	Get user by ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"User ID"
// @Success		200	{object}	models.Response{data=[]models.User} - user retrieved successfully
// @Failure		404	{object}	models.Response{data=[]models.Err} - user not found
// @Failure		410	{object}	models.Response{data=[]models.Err} - user was deleted
// @Failure		503	{object}	models.Response{data=[]models.Err} - no connection to db was established
// @Router			/users/{id} [get]
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

// CreateUser is a function to create a new user
//
// @Summary		Create user
// @Description	Create User
// @Tags			users
// @Accept			json
// @Produce		json
// @Param		set	body		models.User	true	"user"
// @Success		201	{object}	models.Response{data=[]models.Err}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/ [post]
func (h *Handler) Login(c *fiber.Ctx) error {

	//	@Failure	401	{object}	models.Response{}
	//parse auth header
	auth = h.controller.CAuthenticate(authHeader)
	if !auth.data.success {
		return c.Status(auth.StatusCode).JSON(auth)
	}

	response = h.controller.CLoginUser(auth.data)
	if !auth.data.success {
		return c.Status(auth.StatusCode).JSON(auth)
	}

	utils.ErrorUnauthorized("test")




	getCurrentUserFromAuth(auth.data)

	var user *models.User
	var response *models.Response
	if err := c.BodyParser(&user);
	err != nil {
		//	@Failure	400	{object}	models.Response{}
		return c.Status(http.StatusNotAcceptable).JSON(models.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	response = h.controller.CCreateUser(user)
	return c.status(createdUser.StatusCode).JSON(createdUser)


	//	@Failure	503	{object}	models.Response{}
	// 	if no connection to db was established
	createdUser, err := h.controller.CCreateUser(user)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(utils.Response(
			http.StatusAccepted, 
			fmt.Sprintf("User created su"), 
			nil,
			))
	}
	h.me = createdUser

	//	@Success	201	{object}	models.User{}
	return c.Status(http.StatusCreated).JSON(utils.Response(
		http.StatusCreated, 
		fmt.Sprintf("User created successfully"), 
		nil,
		))
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

	id = h.controller.getIdFromAuthHeader(c)

	user, err := h.controller.CGetUserById(id)

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	// 	@Failure	410	{object}	models.Response{}
	// if user was deleted

	//	@Success	200	{object}	models.Response{}
	if h.me != nil {
		//return c.Response(http.StatusAccepted, fmt.Sprintf("Last saved person is: %s", h.me.Username), nil)
	
		return c.Status(http.StatusAccepted).JSON(utils.Response(
			http.StatusAccepted, 
			fmt.Sprintf("Last saved person is: %s", h.me.Username), 
			nil,
			))
	}
	return c.Status(http.StatusNotFound).JSON(utils.ErrorToResponse(
		http.StatusNotFound, 
		fmt.Sprintf("A current user was not found"), 
		"",
		))
}

// UpdateUser is a function to update the given user from the databse
//
// @Summary		Update given user
// @Description	Update given user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param		set	body		models.User	true	"user"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		403	{object}	models.Response{data=[]models.Err}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/users/{id} [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return c.Status(http.StatusNotAcceptable).JSON(models.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure		400	{object}	models.Response{}
	// something with body is wrong/missing

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	updatedPerson, err := h.controller.CUpdateUser(user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Success	200	{object}	models.User{}
	return c.JSON(models.Response{
		Success: true,
		Message: fmt.Sprintf("User %v updated successfully", updatedPerson.Username),
		Data:    nil,
	})
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

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	deletedUser, err := h.controller.CDeleteUser(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Success	200	{object}	models.Response{data=[]models.User}
	return c.JSON(models.Response{
		Success: true,
		Message: fmt.Sprintf("User %s deleted successfully", deletedUser.Username),
		Data:    nil,
	})
}
