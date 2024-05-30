package errors

import (
	"net/http"
)

type ProductNotFoundError string

// Error represents the error message
func (e ProductNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e ProductNotFoundError) ErrCode() string {
	return ProductNotFound
}

// StatusCode represents the HTTP status code
func (e ProductNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
