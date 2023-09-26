package users_handler

// 410 är bra att använda

import (
	"birdai/src/internal/handlers"

	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP = handlers.ResponseHTTP

// GetUser is a function to get a user by ID
//
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	ResponseHTTP{data=[]users_handler.User}
//	@Failure		404	{object}	ResponseHTTP{}
// 	@Failure		410	{object}	ResponseHTTP{}
//	@Failure		503	{object}	ResponseHTTP{}
//	@Router			/users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	//	@Failure	503	{object}	ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	ResponseHTTP{}
	// if user not found

	// 	@Failure	410	{object}	ResponseHTTP{}
	// if user was deleted

	//	@Success	200	{object}	ResponseHTTP{data=[]users_handler.User}
	return c.JSON(handlers.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User with id %v found. (not implemented)", id),
		Data:    id,
	})
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
//	@Success		200	{object}	ResponseHTTP{data=[]users_handler.User}
//	@Failure		401	{object}	ResponseHTTP{}
//	@Failure		503	{object}	ResponseHTTP{}
//	@Router			/users/list [get]
func ListUsers(c *fiber.Ctx) error {
	queries := c.Queries()
	set := queries["set"]
	search := queries["search"]

	//	@Failure	401	{object}	ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	ResponseHTTP{}
	// if user not found

	return c.JSON(handlers.ResponseHTTP{
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
//	@Success		201	{object}	ResponseHTTP{}
//	@Failure		400	{object}	ResponseHTTP{}
//	@Failure		401	{object}	ResponseHTTP{}
//	@Failure		503	{object}	ResponseHTTP{}
//	@Router			/users/ [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(TokenUser)
	if err := c.BodyParser(&user); err != nil {
		//	@Failure	400	{object}	ResponseHTTP{}
		return c.Status(http.StatusNotAcceptable).JSON(handlers.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//	@Failure	401	{object}	ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	ResponseHTTP{}
	// 	if no connection to db was established

	//	@Success	201	{object}	ResponseHTTP{}
	return c.Status(http.StatusCreated).JSON(handlers.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User %v created sucessfully. (not implemented) ", user.User.Username),
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
//	@Success		200	{object}	ResponseHTTP{data=[]users_handler.User}
//	@Failure		401	{object}	ResponseHTTP{}
//	@Failure		404	{object}	ResponseHTTP{}
// 	@Failure		410	{object}	ResponseHTTP{}
//	@Failure		503	{object}	ResponseHTTP{}
//	@Router			/users/me [get]
func GetUserMe(c *fiber.Ctx) error {

	//	@Failure	401	{object}	ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	ResponseHTTP{}
	// if user not found

	// 	@Failure	410	{object}	ResponseHTTP{}
	// if user was deleted

	//	@Success	200	{object}	ResponseHTTP{data=[]users_handler.User}
	return c.JSON(handlers.ResponseHTTP{
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
//	@Success		200	{object}	ResponseHTTP{}
//	@Failure		400	{object}	ResponseHTTP{}
//	@Failure		401	{object}	ResponseHTTP{}
//	@Failure		403	{object}	ResponseHTTP{}
// 	@Failure		404	{object}	ResponseHTTP{}
//	@Failure		503	{object}	ResponseHTTP{}
//	@Router			/users/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	//	@Failure	401	{object}	ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	ResponseHTTP{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure		400	{object}	ResponseHTTP{}
	// something with body is wrong/missing

	//	@Failure	503	{object}	ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	ResponseHTTP{}
	// if user not found

	//	@Success	200	{object}	ResponseHTTP{data=[]users_handler.User}
	return c.JSON(handlers.ResponseHTTP{
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
//	@Success		200	{object}	ResponseHTTP{}
//	@Failure		401	{object}	ResponseHTTP{}
//	@Failure		403	{object}	ResponseHTTP{}
// 	@Failure		404	{object}	ResponseHTTP{}
//	@Failure		503	{object}	ResponseHTTP{}
//	@Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	//	@Failure	401	{object}	ResponseHTTP{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	ResponseHTTP{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure	503	{object}	ResponseHTTP{}
	// if no connection to db was established

	//	@Failure	404	{object}	ResponseHTTP{}
	// if user not found

	//	@Success	200	{object}	ResponseHTTP{data=[]users_handler.User}
	return c.JSON(handlers.ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User %v deleted successfully. not implemented", id),
		Data:    nil,
	})
}
