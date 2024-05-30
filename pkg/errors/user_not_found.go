package errors

import (
	"net/http"
)

type UserNotFoundError string

// Error represents the error message
func (e UserNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e UserNotFoundError) ErrCode() string {
	return UserNotFound
}

// StatusCode represents the HTTP status code
func (e UserNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
