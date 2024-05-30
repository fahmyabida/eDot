package domain

import (
	"context"
	"time"
)

type User struct {
	ID          string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	PhoneNumber string     `json:"phone_number"`
	Password    string     `json:"password"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type UserLoginRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	UserId  string `json:"user_id,omitempty"`
}

type IUserUsecase interface {
	Login(ctx context.Context, payload *UserLoginRequest) (UserLoginResponse, error)
	Register(ctx context.Context, payload *User) error
}

type IUserRepo interface {
	FindByEmailOrPhoneNumber(ctx context.Context, username, phoneNumber string) (user User, err error)
	Create(ctx context.Context, user *User) error
}
