package domain

import (
	"context"
	"time"
)

type Product struct {
	ID                string     `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	Price             float64    `json:"price"`
	AvailableQuantity int        `json:"available_quantity"`
	ReservedQuantity  int        `json:"reserved_quantity"`
	SoldQuantity      int        `json:"sold_quantity"`
	CreatedAt         *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}

type ProductGetAllResponse struct {
	Data  []Product `json:"data"`
	Total int64     `json:"total"`
}

type GetAllProductsPayload struct {
	ID          string `json:"-" query:"id"`
	Name        string `json:"-" query:"name"`
	Description string `json:"-" query:"description"`
	SortBy      string `json:"-" query:"sort_by"`
	Limit       int    `json:"-" query:"limit"`
	Offset      int    `json:"-" query:"offset"`
}

type IProductRepo interface {
	UpdateStockWarehouse(ctx context.Context, active bool, warehouseStocks []WarehouseStock) error
	UpdateStockDeducted(context.Context, *[]OrderProduct) (float64, error)
	FindByID(context.Context, string) (Product, error)
	GetAll(context.Context, *GetAllProductsPayload) ([]Product, int64, error)
	UpdateStockRevert(context.Context, *[]OrderItem) error
}

type IProductUsecase interface {
	GetAllProducts(context.Context, *GetAllProductsPayload) (ProductGetAllResponse, error)
}
