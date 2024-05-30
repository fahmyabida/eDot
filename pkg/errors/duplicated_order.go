package errors

import (
	"net/http"
)

type DuplicateOrderError string

// Error represents the error message
func (e DuplicateOrderError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e DuplicateOrderError) ErrCode() string {
	return DuplicateOrder
}

// StatusCode represents the HTTP status code
func (e DuplicateOrderError) StatusCode() int {
	return http.StatusConflict
}
