package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Healthcheck struct {
	Status       string       `json:"status"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

type Dependency struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// InitHealthcheckHandler will initialize the HTTP handlers to perform health check.
func InitHealthcheckHandler(e *echo.Group) {
	e.GET("/liveness", checkLiveness)
	e.GET("/readiness", checkReadiness)
}

func checkLiveness(c echo.Context) (err error) {
	result := Healthcheck{Status: http.StatusText(http.StatusOK)}
	return c.JSON(http.StatusOK, result)
}

func checkReadiness(c echo.Context) (err error) {
	result := Healthcheck{Status: http.StatusText(http.StatusOK)}
	return c.JSON(http.StatusOK, result)
}
