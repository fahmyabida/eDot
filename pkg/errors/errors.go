package errors

// ServiceError represent the contract for all service errors.
type ServiceError interface {
	// Error to comply with the error interface.
	Error() string
	// ErrCode returns the error code based on the error type.
	ErrCode() string
	// StatusCode returns the HTTP status code based on the error type.
	StatusCode() int
}
