package repository

import (
	"context"
	"fmt"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"gorm.io/gorm"
)

// WarehouseRepository ...
type WarehouseRepository struct {
	DB *gorm.DB
}

// NewWarehouseRepository will return
func NewWarehouseRepository(db *gorm.DB) domain.IWarehouseRepo {
	return &WarehouseRepository{
		DB: db,
	}
}

func (r WarehouseRepository) FindByID(ctx context.Context, ID string) (data domain.Warehouse, err error) {
	dbResult := r.DB.Model(&data).Where("id = ?", ID).Find(&data)
	if dbResult.RowsAffected == 0 {
		err = pkgErrors.WarehouseNotFoundError(fmt.Sprintf(pkgErrors.ErrWarehouseNotFoundWithID, ID))
		return
	}

	return data, err

}

func (r WarehouseRepository) Update(ctx context.Context, data *domain.Warehouse) (err error) {
	dbResult := r.DB.Model(data).Where("id = ?", data.ID).Updates(data)
	if err := dbResult.Error; err != nil {
		return err
	}
	return nil
}
