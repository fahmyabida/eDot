package repository

import (
	"context"
	"fmt"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository will return
func NewUserRepository(db *gorm.DB) domain.IUserRepo {
	return &UserRepository{
		DB: db,
	}
}

func (r UserRepository) FindByEmailOrPhoneNumber(ctx context.Context, email, phoneNumber string) (user domain.User, err error) {

	dbResult := r.DB.Model(&user).Where("email = ? OR phone_number = ?", email, phoneNumber).Find(&user)
	if dbResult.RowsAffected == 0 {
		err = pkgErrors.UserNotFoundError(fmt.Sprintf(pkgErrors.ErrUserNotFound, email, phoneNumber))
		return
	}

	return user, err
}

func (r UserRepository) Create(ctx context.Context, data *domain.User) (err error) {
	dbResult := r.DB.WithContext(ctx).Create(data)
	if dbResult.Error != nil {
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		postgresError, ok := dbResult.Error.(*pgconn.PgError)
		if ok && postgresError.Code == "23505" {
			return pkgErrors.DuplicateUserError(pkgErrors.ErrDuplicateUser)
		}
		return dbResult.Error
	}

	return
}
