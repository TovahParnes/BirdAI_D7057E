package users

import (

	//"swagger/database"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Token string `json:"token" xml:"token" form:"token"`
    Id string `json:"_id" xml:"_id" form:"_id"`
    Username string `json:"username" xml:"username" form:"username"`
}

type Token struct {
    Token string `json:"token" xml:"token" form:"token"`
}

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// GetUser is a function to get a user by ID
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} ResponseHTTP{data=[]models.User}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	// @Failure 503 {object} ResponseHTTP{}
	// if no connection to db was established

	// @Failure 404 {object} ResponseHTTP{}
	// if user not found

	// @Success 200 {object} ResponseHTTP{data=[]models.User}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User with id %v found. (not implemented)", id),
		Data:    id,
	})
}

// GetAllUsers is a function to get a set of all users from database
// @Summary Get set of all users
// @Description Get set of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.User}
// @Failure 401 {object} ResponseHTTP{}
// @Failure 406 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /users/set={set} [get]
func GetAllUsers(c *fiber.Ctx) error {
	set := c.Params("set")

	// @Failure 401 {object} ResponseHTTP{}
	// Authenticate(body.token)

	// @Failure 503 {object} ResponseHTTP{}
	// if no connection to db was established

	// @Failure 404 {object} ResponseHTTP{}
	// if user not found

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("Users from set %v found.", set),
		Data:    set,
	})
}

// CreateUser is a function to create a new user
// @Summary Create user
// @Description Create User
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {object} ResponseHTTP{}
// @Failure 401 {object} ResponseHTTP{}
// @Failure 406 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /users/ [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(&user); err != nil {
		// @Failure 406 {object} ResponseHTTP{}
		return c.Status(http.StatusNotAcceptable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// @Failure 401 {object} ResponseHTTP{}
	// Authenticate(body.token)

	// @Failure 503 {object} ResponseHTTP{}
	// if no connection to db was established
	
	// @Success 201 {object} ResponseHTTP{}
	return c.Status(http.StatusCreated).JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User %v created sucessfully. (not implementeed) ", user.Username),
		Data:    user,
	})
}

// GetUserMe is a function to get the current user from the databse
// @Summary Get current user
// @Description Get current user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.User}
// @Failure 401 {object} ResponseHTTP{}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /users/ [post]
func GetUserMe(c *fiber.Ctx) error {
	token := new(Token)
	if err := c.BodyParser(&token); err != nil {
		// @Failure 406 {object} ResponseHTTP{}
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// @Failure 401 {object} ResponseHTTP{}
	// Authenticate(body.token)

	// @Failure 503 {object} ResponseHTTP{}
	// if no connection to db was established

	// @Failure 404 {object} ResponseHTTP{}
	// if user not found
	
	// @Success 200 {object} ResponseHTTP{data=[]models.User}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("I am user %v.", token.Token),
		Data:    token,
	})
}