package handlers

// 410 är bra att använda

import (
	"birdai/src/internal/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetUserById is a function to get a user by ID
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.ResponseHTTP{data=[]models.User}
//	@Failure		404	{object}	models.ResponseHTTP{}
//	@Failure		410	{object}	models.ResponseHTTP{}
//	@Failure		503	{object}	models.ResponseHTTP{}
//	@Router			/users/{id} [get]
func (h *Handler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.controller.CGetUserById(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Failure	503	{object}	models.ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.ResponseHTTP{}
	// if user not found

	// 	@Failure	410	{object}	models.ResponseHTTP{}
	// if user was deleted

	//	@Success	200	{object}	models.ResponseHTTP{data=[]models.User}
	return c.JSON(user)
}

// ListUsers is a function to get a set of all users from database
//
//	@Summary		List all users of a specified set
//	@Description	List all users of a specified set
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			set	query		int	false	"Set of users"
//	@Param			search	query	string	false	"Search parameter for user"
//	@Success		200	{object}	models.ResponseHTTP{data=[]models.User}
//	@Failure		401	{object}	models.ResponseHTTP{}
//	@Failure		503	{object}	models.ResponseHTTP{}
//	@Router			/users/list [get]
func (h *Handler) ListUsers(c *fiber.Ctx) error {
	queries := c.Queries()
	set := queries["set"]
	search := queries["search"]

	//	@Failure	401	{object}	models.ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.ResponseHTTP{}
	// if user not found

	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("Users from set %v found. Search param %v", set, search),
		Data:    set,
	})
}

// CreateUser is a function to create a new user
//
//	@Summary		Create user
//	@Description	Create User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	models.ResponseHTTP{}
//	@Failure		400	{object}	models.ResponseHTTP{}
//	@Failure		401	{object}	models.ResponseHTTP{}
//	@Failure		503	{object}	models.ResponseHTTP{}
//	@Router			/users/ [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		//	@Failure	400	{object}	models.ResponseHTTP{}
		return c.Status(http.StatusNotAcceptable).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Failure	401	{object}	models.ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.ResponseHTTP{}
	// 	if no connection to db was established
	err := h.controller.CCreateUser(user)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Success	201	{object}	models.ResponseHTTP{}
	return c.Status(http.StatusCreated).JSON(models.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User %v created sucessfully ", user.Username),
		Data:    user,
	})
}

// GetUserMe is a function to get the current user from the databse
//
//	@Summary		Get current user
//	@Description	Get current user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ResponseHTTP{data=[]models.User}
//	@Failure		401	{object}	models.ResponseHTTP{}
//	@Failure		404	{object}	models.ResponseHTTP{}
//	@Failure		410	{object}	models.ResponseHTTP{}
//	@Failure		503	{object}	models.ResponseHTTP{}
//	@Router			/users/me [get]
func (h *Handler) GetUserMe(c *fiber.Ctx) error {

	//	@Failure	401	{object}	models.ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.ResponseHTTP{}
	// if user not found

	// 	@Failure	410	{object}	models.ResponseHTTP{}
	// if user was deleted

	//	@Success	200	{object}	models.ResponseHTTP{data=[]models.User}
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("I am user -me-. not implemented"),
		Data:    nil,
	})
}

// UpdateUser is a function to update the given user from the databse
//
//	@Summary		Update given user
//	@Description	Update given user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.ResponseHTTP{}
//	@Failure		400	{object}	models.ResponseHTTP{}
//	@Failure		401	{object}	models.ResponseHTTP{}
//	@Failure		403	{object}	models.ResponseHTTP{}
//	@Failure		404	{object}	models.ResponseHTTP{}
//	@Failure		503	{object}	models.ResponseHTTP{}
//	@Router			/users/{id} [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	//	@Failure	401	{object}	models.ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	models.ResponseHTTP{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure		400	{object}	models.ResponseHTTP{}
	// something with body is wrong/missing

	//	@Failure	503	{object}	models.ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.ResponseHTTP{}
	// if user not found

	//	@Success	200	{object}	models.ResponseHTTP{data=[]models.User}
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User %v updated successfully. not implemented", id),
		Data:    nil,
	})
}

// DeleteUser is a function to update the given user from the databse
//
//	@Summary		Delete given user
//	@Description	Delete given user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.ResponseHTTP{}
//	@Failure		401	{object}	models.ResponseHTTP{}
//	@Failure		403	{object}	models.ResponseHTTP{}
//	@Failure		404	{object}	models.ResponseHTTP{}
//	@Failure		503	{object}	models.ResponseHTTP{}
//	@Router			/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	//	@Failure	401	{object}	models.ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	models.ResponseHTTP{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure	503	{object}	models.ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.ResponseHTTP{}
	// if user not found

	//	@Success	200	{object}	models.ResponseHTTP{data=[]models.User}
	return c.JSON(models.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User %v deleted successfully. not implemented", id),
		Data:    nil,
	})
}
