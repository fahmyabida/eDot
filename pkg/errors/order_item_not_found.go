package errors

import (
	"net/http"
)

type OrderItemNotFoundError string

// Error represents the error message
func (e OrderItemNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e OrderItemNotFoundError) ErrCode() string {
	return OrderItemNotFound
}

// StatusCode represents the HTTP status code
func (e OrderItemNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
