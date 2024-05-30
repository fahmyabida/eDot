package repository

import (
	"context"
	"time"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// OrderRepository ...
type OrderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository will return
func NewOrderRepository(db *gorm.DB) domain.IOrderRepo {
	return &OrderRepository{
		DB: db,
	}
}

func (r OrderRepository) Create(ctx context.Context, data *domain.Order) (err error) {
	dbResult := r.DB.WithContext(ctx).Create(data)
	if dbResult.Error != nil {
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		postgresError, ok := dbResult.Error.(*pgconn.PgError)
		if ok && postgresError.Code == "23505" {
			return pkgErrors.DuplicateOrderError(pkgErrors.ErrDuplicateOrder)
		}
		return dbResult.Error
	}

	return
}

func (r OrderRepository) FindByID(ctx context.Context, id string) (order domain.Order, err error) {
	dbResult := r.DB.Model(&order).Where("id = ?", id).Find(&order)
	if dbResult.RowsAffected == 0 {
		return order, pkgErrors.OrderNotFoundError(pkgErrors.ErrOrderNotFound)
	}
	return order, dbResult.Error
}

func (r OrderRepository) GetOrdersByStatusAndCreatedAtLessThan(ctx context.Context, status string, createdAt time.Time) (orders []domain.Order, err error) {
	dbResult := r.DB.Model(&orders).Where("status = ? AND created_at < ?", status, createdAt).Find(&orders)
	if dbResult.RowsAffected == 0 {
		return orders, pkgErrors.OrderNotFoundError(pkgErrors.ErrOrderNotFound)
	} else if err = dbResult.Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func (r OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	dbResult := r.DB.Model(order).Where("id = ?", order.ID).Updates(order)
	if err := dbResult.Error; err != nil {
		return err
	}
	return nil
}
