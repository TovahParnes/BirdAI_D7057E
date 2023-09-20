package users

import (

	//"swagger/database"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
    Id string `json:"_id" xml:"_id" form:"_id"`
    Username string `json:"username" xml:"username" form:"username"`
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
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User with id %v found.", id),
		Data:    id,
	})
/*
	db := database.DBConn

	book := new(models.Book)
	if err := db.First(&book, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Book with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get book by ID.",
		Data:    *book,
	})
	*/
}

// GetAllUsers is a function to get a set of all users from database
// @Summary Get set of all users
// @Description Get set of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.User}
// @Failure 503 {object} ResponseHTTP{}
// @Router /users/set={set} [get]
func GetAllUsers(c *fiber.Ctx) error {
	set := c.Params("set")
	/*
	db := database.DBConn

	var books []models.Book
	if res := db.Find(&books); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get all books.",
		Data:    books,
	})
*/
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("Users from set %v found.", set),
		Data:    set,
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
// @Router /users/me [post]
func GetUserMe(c *fiber.Ctx) error {

	
	body := new(User)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("I am user %v.", body.Username),
		Data:    body,
	})
}