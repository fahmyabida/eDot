package repository

import (
	"context"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// OrderItemRepository ...
type OrderItemRepository struct {
	DB *gorm.DB
}

// NewOrderItemRepository will return
func NewOrderItemRepository(db *gorm.DB) domain.IOrderItemRepo {
	return &OrderItemRepository{
		DB: db,
	}
}
func (r OrderItemRepository) Create(ctx context.Context, data *[]domain.OrderItem) (err error) {
	dbResult := r.DB.WithContext(ctx).Create(data)
	if dbResult.Error != nil {
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		postgresError, ok := dbResult.Error.(*pgconn.PgError)
		if ok && postgresError.Code == "23505" {
			return pkgErrors.DuplicateOrderItemError(pkgErrors.ErrDuplicateOrderItem)
		}
		return dbResult.Error
	}

	return
}

func (r OrderItemRepository) GetOrderItemsByOrderId(ctx context.Context, orderId string) (orderItems []domain.OrderItem, err error) {
	dbResult := r.DB.WithContext(ctx).Model(&orderItems).Where("order_id = ?", orderId).Find(&orderItems)
	if dbResult.RowsAffected == 0 {
		return orderItems, pkgErrors.OrderItemNotFoundError(pkgErrors.ErrOrderItemNotFound)
	}
	return orderItems, dbResult.Error
}
