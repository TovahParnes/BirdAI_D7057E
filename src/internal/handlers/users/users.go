package users

import (

	//"swagger/database"
	//"swagger/models"

	"github.com/gofiber/fiber/v2"
)

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
// @Success 200 {object} ResponseHTTP{data=[]models.Book}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
/*
	id := c.Params("id")
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
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get user.",
		Data:    "test_User",
	})
}