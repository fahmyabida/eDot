package domain

import (
	"context"
	"time"
)

type OrderItem struct {
	ID        string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrderId   string     `json:"order_id"`
	ProductId string     `json:"product_id"`
	Quantity  int        `json:"quantity"`
	UnitPrice float64    `json:"unit_price"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type IOrderItemRepo interface {
	Create(context.Context, *[]OrderItem) error
	GetOrderItemsByOrderId(context.Context, string) ([]OrderItem, error)
}
