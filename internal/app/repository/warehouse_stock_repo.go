package repository

import (
	"context"
	"fmt"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// WarehouseStockRepository ...
type WarehouseStockRepository struct {
	DB *gorm.DB
}

// NewWarehouseStockRepository will return
func NewWarehouseStockRepository(db *gorm.DB) domain.IWarehouseStockRepo {
	return &WarehouseStockRepository{
		DB: db,
	}
}

func (r WarehouseStockRepository) TransferStock(ctx context.Context, transferMode domain.TransferStockMode, payloads []domain.WarehouseStock) (err error) {
	switch transferMode {
	case domain.SOURCE:
		for _, rowPayload := range payloads {
			var warehouseStock = domain.WarehouseStock{}
			dbResult := r.DB.Model(&warehouseStock).Where("product_id = ? AND warehouse_id = ?", rowPayload.ProductId, rowPayload.WarehouseId).Find(&warehouseStock)
			if err = dbResult.Error; err != nil {
				return err
			}
			warehouseStock.Quantity -= rowPayload.Quantity
			dbResult = r.DB.Model(&warehouseStock).Where("product_id = ? AND warehouse_id = ?", rowPayload.ProductId, rowPayload.WarehouseId).Updates(&warehouseStock)
			if err = dbResult.Error; err != nil {
				return err
			}
		}
		return nil
	case domain.DESTINATION:
		for _, rowPayload := range payloads {
			var warehouseStock = domain.WarehouseStock{}
			r.DB.Model(&warehouseStock).Where("product_id = ? AND warehouse_id = ?", rowPayload.ProductId, rowPayload.WarehouseId).Find(&warehouseStock)
			if warehouseStock.ID == "" {
				r.DB.Create(&rowPayload)
				continue
			}
			warehouseStock.Quantity += rowPayload.Quantity
			dbResult := r.DB.Model(&warehouseStock).Where("product_id = ? AND warehouse_id = ?", rowPayload.ProductId, rowPayload.WarehouseId).Updates(&warehouseStock)
			if err = dbResult.Error; err != nil {
				return err
			}
		}

	}
	return nil
}

func (r WarehouseStockRepository) GetByWarehouseIAndProductIds(ctx context.Context, warehouseID string, productIDs []string) (datas []domain.WarehouseStock, err error) {
	dbResult := r.DB.Model(&datas).Where("warehouse_id = ? AND product_id IN (?)", warehouseID, productIDs).Find(&datas)
	if err = dbResult.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = pkgErrors.WarehouseStockNotFoundError(pkgErrors.ErrWarehouseStockNotFound)
		}
		return datas, err
	}
	return datas, err
}

func (r WarehouseStockRepository) GetAll(ctx context.Context, params *domain.GetAllWarehouseStocksPayload) (warehouseStocks []domain.WarehouseStock, count int64, err error) {
	db := r.DB.WithContext(ctx).Limit(10).Offset(0)

	db, err = r.setQueryForGetAll(db, params)
	if err != nil {
		return
	}

	db = db.Find(&warehouseStocks).Offset(-1).Limit(-1).Count(&count) // reset offset & limit for the count
	if err = db.Error; err != nil {
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		postgresError, ok := db.Error.(*pgconn.PgError)
		if ok && postgresError.Code == "42703" {
			err = pkgErrors.InvalidColumnError(postgresError.Message)
		} else if err == gorm.ErrRecordNotFound {
			err = pkgErrors.WarehouseStockNotFoundError(pkgErrors.ErrWarehouseStockNotFound)
		}
		return
	}

	return warehouseStocks, count, err
}

func (r WarehouseStockRepository) setQueryForGetAll(db *gorm.DB, params *domain.GetAllWarehouseStocksPayload) (queryDB *gorm.DB, err error) {

	if params != nil {
		db = r.setLimitAndOffsetGetAll(db, params.Limit, params.Offset)
		db = r.setConditionForGetAll(db, "id", params.ID)
		db = r.setConditionForGetAll(db, "warehouse_id", params.WarehouseID)
		db = r.setConditionForGetAll(db, "product_id", params.ProductId)

		if params.SortBy != "" {
			db, err = ApplySortByQuery(db, params.SortBy)
			if err != nil {
				return
			}
		} else {
			db = db.Order("created_at DESC")
		}
	} else {
		db = db.Order("created_at DESC")
	}
	queryDB = db
	return
}

func (r WarehouseStockRepository) setLimitAndOffsetGetAll(db *gorm.DB, limit, offset int) *gorm.DB {
	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}
	return db
}

func (r WarehouseStockRepository) setConditionForGetAll(db *gorm.DB, field, criteria string) *gorm.DB {
	if criteria != "" {
		query := fmt.Sprintf("%s = '%s'", field, criteria)
		db = db.Where(query)
	}
	return db
}
