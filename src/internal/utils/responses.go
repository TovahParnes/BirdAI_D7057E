package utils

import (
	"birdai/src/internal/models"
	"net/http"
	"time"
)


func ResponseHTTP(statusCode int, message string, data interface{}) models.ResponseHTTP {
	return models.ResponseHTTP{
		Success: statusCode >= 200 && statusCode <= 299,
		StatusCode : statusCode,
		StatusName: http.StatusText(statusCode),
		Data: data,
		Timestamp: time.Now().Format(time.RFC3339),
		Message: message,
		Description: "",
	}
}


func ErrorToResponseHTTP(statusCode int, message string, err error) models.ResponseHTTP {
return models.ResponseHTTP{
Success: statusCode >= 200 && statusCode <= 299,
StatusCode : statusCode,
StatusName: http.StatusText(statusCode),
Data: nil,
Timestamp: time.Now().Format(time.RFC3339),
Message: message,
Description: err.Error(),
}
}
