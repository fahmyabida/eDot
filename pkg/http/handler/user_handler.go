package handler

import (
	"net/http"

	"github.com/fahmyabida/eDot/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase domain.IUserUsecase
}

func InitUserHandler(e *echo.Group, eNoAuth *echo.Group, userUsecase domain.IUserUsecase) {
	handler := UserHandler{UserUsecase: userUsecase}

	eNoAuth.POST("/user/login", handler.UserLoginHandler)
	e.POST("/user", handler.RegisterUserHandler)
}

func (h *UserHandler) UserLoginHandler(c echo.Context) error {

	var user domain.UserLoginRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	ctx := c.Request().Context()

	response, err := h.UserUsecase.Login(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) RegisterUserHandler(c echo.Context) error {

	var user domain.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	ctx := c.Request().Context()

	err := h.UserUsecase.Register(ctx, &user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}
