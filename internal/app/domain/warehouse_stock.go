package domain

import (
	"context"
	"time"
)

type TransferStockMode string

const (
	SOURCE      TransferStockMode = "SOURCE"
	DESTINATION TransferStockMode = "DESTINATION"
)

type WarehouseStock struct {
	ID          string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	WarehouseId string     `json:"warehouse_id"`
	ProductId   string     `json:"product_id"`
	Quantity    int        `json:"quantity"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type GetAllWarehouseStocksPayload struct {
	ID          string `json:"-" query:"id"`
	WarehouseID string `json:"-" query:"warehouse_id"`
	ProductId   string `json:"-" query:"product_id"`
	SortBy      string `json:"-" query:"sort_by"`
	Limit       int    `json:"-" query:"limit"`
	Offset      int    `json:"-" query:"offset"`
}

type IWarehouseStockRepo interface {
	TransferStock(ctx context.Context, transferMode TransferStockMode, payloads []WarehouseStock) (err error)
	GetAll(context.Context, *GetAllWarehouseStocksPayload) ([]WarehouseStock, int64, error)
	GetByWarehouseIAndProductIds(context.Context, string, []string) ([]WarehouseStock, error)
}
