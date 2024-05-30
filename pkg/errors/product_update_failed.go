package errors

import (
	"net/http"
)

type ProductUpdateFailedError string

// Error represents the error message
func (e ProductUpdateFailedError) Error() string {
	return string(e)
}

// ErrCode represents the Xendit error code
func (e ProductUpdateFailedError) ErrCode() string {
	return ProductUpdateFailed
}

// StatusCode represents the HTTP status code
func (e ProductUpdateFailedError) StatusCode() int {
	return http.StatusConflict
}
