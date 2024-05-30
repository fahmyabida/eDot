package errors

import (
	"net/http"
)

type InvalidAmountError string

// Error represents the error message
func (e InvalidAmountError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e InvalidAmountError) ErrCode() string {
	return InvalidAmount
}

// StatusCode represents the HTTP status code
func (e InvalidAmountError) StatusCode() int {
	return http.StatusNotFound
}
