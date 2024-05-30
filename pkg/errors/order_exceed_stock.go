package errors

import (
	"net/http"
)

type OrderExceedStockError string

// Error represents the error message
func (e OrderExceedStockError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e OrderExceedStockError) ErrCode() string {
	return OrderExceedStock
}

// StatusCode represents the HTTP status code
func (e OrderExceedStockError) StatusCode() int {
	return http.StatusBadRequest
}
