package utils

import (
	"birdai/src/internal/models"
	"net/http"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
)


func CreationResponseToStatus(c *fiber.Ctx, response models.Response) error {
	return c.Status(http.StatusCreated).JSON(response)
}

func ResponseToStatus(c *fiber.Ctx, response models.Response) error {
	if IsTypeError(response){
		return c.Status(response.Data.(models.Err).StatusCode).JSON(response)
	}
	return c.Status(http.StatusOK).JSON(response)
}



func Response(data interface{}) models.Response {
	return models.Response{
		Data: data,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func ErrorToResponse(statusCode int, message string, err string) models.Response {
	return  Response(models.Err{
		StatusCode : statusCode,
		StatusName: http.StatusText(statusCode),
		Message: message,
		Description: err,
	})
}

func IsTypeError(response models.Response) bool {
	return reflect.TypeOf(response.Data) == reflect.TypeOf(models.Err{})
}

func IsType(response models.Response, data interface{}) bool {
	return (reflect.TypeOf(response.Data) == reflect.TypeOf(data)) ||
		(reflect.TypeOf(response.Data) == reflect.PtrTo(reflect.TypeOf(data)))


}



// if the JTW token is invalid
func ErrorUnauthorized(err string) models.Response{
	return ErrorToResponse(http.StatusUnauthorized, "Could not authorize user", err)
}

func ErrorForbidden(err string) models.Response{
	return ErrorToResponse(http.StatusForbidden, "User does not have access to this request", err)
}

func ErrorNotFoundInDatabase(err string) models.Response{
	return ErrorToResponse(http.StatusNotFound, "Could not find any document in the database", err)
}

func ErrorDeleted(err string) models.Response{
	return ErrorToResponse(http.StatusGone, "The document with the given ID has been deleted", err)
}

func ErrorParams(err string) models.Response{
	return ErrorToResponse(http.StatusBadRequest, "Could not parse parameters", err)
}

func ErrorCollectionNotFound(err string) models.Response{
	return ErrorToResponse(http.StatusNotFound, "Could not find collection", err)
}

func ErrorNotImplemented(err string) models.Response{
	return ErrorToResponse(http.StatusNotImplemented, "Not implemented yet", err)
}