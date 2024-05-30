package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	errors "github.com/fahmyabida/eDot/pkg/errors"
	"github.com/labstack/echo/v4"
)

// DefaultErrorMessage is the default server error message.
const DefaultErrorMessage = "An unexpected error occurred, our team has been notified and will troubleshoot the issue."

// ErrorMiddleware is an echo middleware that handles the outgoing errors.
func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			return errorHandler(c, err)
		}
	}
}

func errorHandler(c echo.Context, err error) error {
	switch e := err.(type) {
	case errors.ServiceError:
		return handledError(c, e)
	case *echo.HTTPError:
		return handledEchoError(c, e)
	default:
		return unhandledError(c, err)
	}
}

func handledEchoError(c echo.Context, err *echo.HTTPError) error {
	errCode := fmt.Sprint(err.Code)
	message := err.Message

	if err.Code == http.StatusInternalServerError {
		log.Default().Println(c.Request().Context(), map[string]interface{}{
			"method":        c.Request().Method,
			"uri":           c.Request().RequestURI,
			"error_code":    errCode,
			"error_message": message,
		}, "http error", err)

		errCode = httpStatusCodeToErrorCode(http.StatusInternalServerError)
		message = DefaultErrorMessage
	}

	return c.JSON(err.Code, map[string]interface{}{
		"error_code": errCode,
		"message":    message,
	})
}

func handledError(c echo.Context, err errors.ServiceError) error {
	errCode := err.ErrCode()
	message := err.Error()

	if err.StatusCode() == http.StatusInternalServerError {
		log.Default().Println(c.Request().Context(), map[string]interface{}{
			"method":        c.Request().Method,
			"uri":           c.Request().RequestURI,
			"error_code":    errCode,
			"error_message": message,
		}, "http error", err)

		errCode = httpStatusCodeToErrorCode(http.StatusInternalServerError)
		message = DefaultErrorMessage
	}

	return c.JSON(err.StatusCode(), map[string]interface{}{
		"error_code": errCode,
		"message":    message,
	})
}

func unhandledError(c echo.Context, err error) error {
	var errCode string
	message := err.Error()

	// Check if err is a HTTP error (to assign the correct error_code in our logs if it is)
	if httpErr, ok := err.(*echo.HTTPError); ok {
		errCode = httpStatusCodeToErrorCode(httpErr.Code)
	} else {
		errCode = httpStatusCodeToErrorCode(http.StatusInternalServerError)
	}

	log.Default().Println(c.Request().Context(), map[string]interface{}{
		"method":        c.Request().Method,
		"uri":           c.Request().RequestURI,
		"error_code":    errCode,
		"error_message": message,
	}, "http error", err)

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error_code": httpStatusCodeToErrorCode(http.StatusInternalServerError),
		"message":    DefaultErrorMessage,
	})
}

// httpStatusCodeToErrorCode returns the HTTP error message of the statusCode (uppercase with underscores).
func httpStatusCodeToErrorCode(statusCode int) string {
	if statusCode == http.StatusInternalServerError {
		return "SERVER_ERROR"
	}
	statusText := http.StatusText(statusCode)
	statusText = strings.ReplaceAll(statusText, " ", "_")
	return strings.ToUpper(statusText)
}
