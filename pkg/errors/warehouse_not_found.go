package errors

import (
	"net/http"
)

type WarehouseNotFoundError string

// Error represents the error message
func (e WarehouseNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e WarehouseNotFoundError) ErrCode() string {
	return WarehouseNotFound
}

// StatusCode represents the HTTP status code
func (e WarehouseNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
