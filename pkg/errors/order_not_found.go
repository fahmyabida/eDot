package errors

import (
	"net/http"
)

type OrderNotFoundError string

// Error represents the error message
func (e OrderNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e OrderNotFoundError) ErrCode() string {
	return OrderNotFound
}

// StatusCode represents the HTTP status code
func (e OrderNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
