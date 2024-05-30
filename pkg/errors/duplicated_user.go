package errors

import (
	"net/http"
)

type DuplicateUserError string

// Error represents the error message
func (e DuplicateUserError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e DuplicateUserError) ErrCode() string {
	return DuplicateUser
}

// StatusCode represents the HTTP status code
func (e DuplicateUserError) StatusCode() int {
	return http.StatusConflict
}
