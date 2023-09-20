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

// GetAllBooks is a function to get all books data from database
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Book}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/books [get]
func GetUser(c *fiber.Ctx) error {
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
		Message: "Success get user.",
		Data:    "test_User",
	})
}