package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/fahmyabida/eDot/cmd/config"
	"github.com/fahmyabida/eDot/internal/app/domain"
	"github.com/fahmyabida/eDot/internal/app/usecase"
	"github.com/fahmyabida/eDot/mocks"
)

func TestUserUsecaseImpl_Login(t *testing.T) {
	mockUserRepo := new(mocks.IUserRepo)
	jwtConfig := &config.JWTConfig{
		SecretKey: "testsecret",
	}

	userUsecase := usecase.NewUserUsecase(mockUserRepo, jwtConfig)

	ctx := context.Background()
	payload := &domain.UserLoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	t.Run("successful login", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		mockUser := domain.User{
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}

		mockUserRepo.On("FindByEmailOrPhoneNumber", ctx, "test@example.com", "").Return(mockUser, nil).Once()

		response, err := userUsecase.Login(ctx, payload)

		assert.NoError(t, err)
		assert.Equal(t, "ok", response.Message)
		assert.NotEmpty(t, response.Token)

		token, _ := jwt.Parse(response.Token, func(token *jwt.Token) (interface{}, error) {
			return jwtConfig.SecretKey, nil
		})

		assert.True(t, token.Valid)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		mockUser := domain.User{
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}

		mockUserRepo.On("FindByEmailOrPhoneNumber", ctx, "test@example.com", "").Return(mockUser, nil).Once()

		payload.Password = "wrongpassword"
		response, err := userUsecase.Login(ctx, payload)

		assert.Error(t, err)
		assert.Equal(t, "wrong password: bcrypt: hashedPassword is not the hash of the given password", response.Message)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("email or phone number not found", func(t *testing.T) {
		mockUserRepo.On("FindByEmailOrPhoneNumber", ctx, "test@example.com", "").Return(nil, errors.New("user not found")).Once()

		payload.Password = "password123"
		response, err := userUsecase.Login(ctx, payload)

		assert.Error(t, err)
		assert.Equal(t, "email or phone number not found", response.Message)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_Register(t *testing.T) {
	mockUserRepo := new(mocks.IUserRepo)
	jwtConfig := &config.JWTConfig{
		SecretKey: "testsecret",
	}
	userUsecase := usecase.NewUserUsecase(mockUserRepo, jwtConfig)

	ctx := context.Background()

	t.Run("successful registration", func(t *testing.T) {
		user := &domain.User{
			Email:       "newuser@example.com",
			Username:    "newuser",
			PhoneNumber: "123456789",
			Password:    "password",
		}

		mockUserRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUsecase.Register(ctx, user)
		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("registration failure", func(t *testing.T) {
		user := &domain.User{
			Email:       "newuser@example.com",
			Username:    "newuser",
			PhoneNumber: "123456789",
			Password:    "password",
		}

		mockUserRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).Return(errors.New("creation failed")).Once()

		err := userUsecase.Register(ctx, user)
		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
