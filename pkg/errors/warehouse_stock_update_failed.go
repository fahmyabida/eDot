package errors

import (
	"net/http"
)

type WarehouseStockUpdateFailedError string

// Error represents the error message
func (e WarehouseStockUpdateFailedError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e WarehouseStockUpdateFailedError) ErrCode() string {
	return WarehouseStockUpdateFailed
}

// StatusCode represents the HTTP status code
func (e WarehouseStockUpdateFailedError) StatusCode() int {
	return http.StatusConflict
}
