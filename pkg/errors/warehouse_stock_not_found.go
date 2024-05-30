package errors

import (
	"net/http"
)

type WarehouseStockNotFoundError string

// Error represents the error message
func (e WarehouseStockNotFoundError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e WarehouseStockNotFoundError) ErrCode() string {
	return WarehouseStockNotFound
}

// StatusCode represents the HTTP status code
func (e WarehouseStockNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
