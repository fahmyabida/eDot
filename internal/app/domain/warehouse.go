package domain

import (
	"context"
	"time"
)

type WarehouseStatus string

const (
	ACTIVE   WarehouseStatus = "ACTIVE"
	INACTIVE WarehouseStatus = "INACTIVE"
)

type Warehouse struct {
	ID        string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type SwitchWarehouseRequest struct {
	WarehouseID string          `json:"warehouse_id"`
	Mode        WarehouseStatus `json:"mode"`
}

type SwitchWarehouseResponse struct {
	Message string          `json:"message"`
	Status  WarehouseStatus `json:"status"`
}

type StockTransferRequest struct {
	SourceWarehouseID      string                 `json:"source_warehouse_id"`
	DestinationWarehouseID string                 `json:"destination_warehouse_id"`
	Products               []StockTransferProduct `json:"products"`
}

type StockTransferProduct struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type StockTransferResponse struct {
	Message string          `json:"message"`
	Status  WarehouseStatus `json:"status"`
}

type IWarehouseUsecase interface {
	TransferStockWarehouse(context.Context, *StockTransferRequest) (StockTransferResponse, error)
	SwitchWarehouse(context.Context, *SwitchWarehouseRequest) (SwitchWarehouseResponse, error)
}

type IWarehouseRepo interface {
	FindByID(context.Context, string) (Warehouse, error)
	Update(context.Context, *Warehouse) error
}
