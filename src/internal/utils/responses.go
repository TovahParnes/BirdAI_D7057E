package utils

import (
	"birdai/src/internal/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreationResponseToStatus(c *fiber.Ctx, response models.Response) error {
	return c.Status(http.StatusCreated).JSON(response)
}

func ResponseToStatus(c *fiber.Ctx, response models.Response) error {
	if !response.Data.(models.Err).Success{
		return c.Status(response.Data.(models.Err).StatusCode).JSON(response)
	}
	return c.Status(http.StatusAccepted).JSON(response)
}



func Response(data interface{}) models.Response {
	return models.Response{
		Data: data,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func ErrorToResponse(statusCode int, message string, err string) models.Response {
	return  Response(models.Err{
		Success: false,
		StatusCode : statusCode,
		StatusName: http.StatusText(statusCode),
		Message: message,
		Description: err,
	})
}



// if the JTW token is invalid
func ErrorUnauthorized(err string) models.Response{
	return ErrorToResponse(http.StatusUnauthorized, "Could not authorize user", err)
}

func ErrorForbidden(err string) models.Response{
	return ErrorToResponse(http.StatusForbidden, "User does not have access to this request", err)
}

func ErrorNotFoundInDatabase(err string) models.Response{
	return ErrorToResponse(http.StatusNotFound, "Could not find any document with the given ID", err)
}

func ErrorDeleted(err string) models.Response{
	return ErrorToResponse(http.StatusGone, "The document with the given ID has been deleted", err)
}

func ErrorParams(err string) models.Response{
	return ErrorToResponse(http.StatusBadRequest, "Could not parse parameters", err)
}
