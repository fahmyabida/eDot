package errors

import (
	"net/http"
)

// InvalidColumn represents and invalid/not found column in table
type InvalidColumnError string

// Error represents the error message
func (e InvalidColumnError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e InvalidColumnError) ErrCode() string {
	return InvalidColumn
}

// StatusCode represents the HTTP status code
func (e InvalidColumnError) StatusCode() int {
	return http.StatusBadRequest
}
