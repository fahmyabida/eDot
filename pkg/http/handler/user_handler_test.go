// user_handler_test.go
package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fahmyabida/eDot/internal/app/domain"
	"github.com/fahmyabida/eDot/mocks"
	"github.com/fahmyabida/eDot/pkg/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserLoginHandler(t *testing.T) {
	mockUsecase := new(mocks.IUserUsecase)
	mockUsecase.On("Login", mock.Anything, mock.Anything).Return(domain.UserLoginResponse{}, errors.New("test error")).Once()
	handler := handler.UserHandler{mockUsecase}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(`{"username":"testuser", "password":"testpass"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler.UserLoginHandler(c)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	mockUsecase.AssertExpectations(t)
}

func TestRegisterUserHandler(t *testing.T) {
	mockUsecase := new(mocks.IUserUsecase)
	handler := handler.UserHandler{UserUsecase: mockUsecase}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(`{"username":"testuser", "password":"testpass"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase.On("Register", mock.Anything, mock.Anything).Return(nil)

	handler.RegisterUserHandler(c)
	assert.Equal(t, http.StatusCreated, rec.Code)

	mockUsecase.AssertExpectations(t)
}
