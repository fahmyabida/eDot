package usecase

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/fahmyabida/eDot/cmd/config"
	"github.com/fahmyabida/eDot/internal/app/domain"
	"github.com/golang-jwt/jwt/v4"
)

type UserUsecaseImpl struct {
	userRepo  domain.IUserRepo
	jwtConfig *config.JWTConfig
}

func NewUserUsecase(userRepo domain.IUserRepo, jwtConfig *config.JWTConfig) domain.IUserUsecase {
	return &UserUsecaseImpl{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
	}
}

func (u *UserUsecaseImpl) Login(ctx context.Context, payload *domain.UserLoginRequest) (domain.UserLoginResponse, error) {

	var response domain.UserLoginResponse

	user, err := u.userRepo.FindByEmailOrPhoneNumber(ctx, payload.Email, payload.PhoneNumber)
	if err != nil {
		response.Message = "email or phone number not found"
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		response.Message = fmt.Sprintf("wrong password: %v", err)
		return response, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtConfig.SecretKey))
	if err != nil {
		response.Message = fmt.Sprintf("error signing token: %v", err)
		return response, err
	}

	response.Message = "ok"
	response.Token = tokenString
	response.UserId = user.ID

	return response, err
}

func (u *UserUsecaseImpl) Register(ctx context.Context, payload *domain.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = u.userRepo.Create(ctx, &domain.User{
		Email:       payload.Email,
		Username:    payload.Username,
		PhoneNumber: payload.PhoneNumber,
		Password:    string(hashedPassword),
	})
	return err
}
