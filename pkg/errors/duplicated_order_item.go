package errors

import (
	"net/http"
)

type DuplicateOrderItemError string

// Error represents the error message
func (e DuplicateOrderItemError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e DuplicateOrderItemError) ErrCode() string {
	return DuplicateOrderItem
}

// StatusCode represents the HTTP status code
func (e DuplicateOrderItemError) StatusCode() int {
	return http.StatusConflict
}
