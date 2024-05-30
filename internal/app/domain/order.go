package domain

import (
	"context"
	"time"
)

type OrderStatus string

const (
	ORDERED  OrderStatus = "ORDERED"
	TIMEOUT  OrderStatus = "TIMEOUT"
	PAYED    OrderStatus = "PAYED"
	SHIPPED  OrderStatus = "SHIPPED"
	COMPLETE OrderStatus = "COMPLETE"
)

type Order struct {
	ID          string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      string     `json:"user_id"`
	OrderDate   *time.Time `json:"order_date,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	TotalAmount float64    `json:"total_amount"`
	Status      string     `json:"status" gorm:"type:text"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type CreateOrderRequest struct {
	UserID   string         `json:"user_id"`
	Products []OrderProduct `json:"products"`
}

type OrderProduct struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     float64
}

type CreteOrderResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type IOrderRepo interface {
	Create(context.Context, *Order) error
	FindByID(context.Context, string) (Order, error)
	GetOrdersByStatusAndCreatedAtLessThan(context.Context, string, time.Time) ([]Order, error)
	Update(context.Context, *Order) error
}

type IOrderUsecase interface {
	Create(context.Context, *CreateOrderRequest) (CreteOrderResponse, error)
	CancelExceededOrder(ctx context.Context) (err error)
}
