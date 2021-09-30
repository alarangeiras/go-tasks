package errors

import "net/http"

func NewNotFound() AppError {
	return newBaseError(http.StatusNotFound, "Resource not found")
}
