package repository

import (
	"context"
	"fmt"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// ProductRepository ...
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository will return
func NewProductRepository(db *gorm.DB) domain.IProductRepo {
	return &ProductRepository{
		DB: db,
	}
}

func (r ProductRepository) UpdateStockWarehouse(ctx context.Context, active bool, warehouseStocks []domain.WarehouseStock) (err error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	mapWarehouseStock := map[string]domain.WarehouseStock{}
	productsIDs := []string{}
	for _, rowWarehouseStocks := range warehouseStocks {
		mapWarehouseStock[rowWarehouseStocks.ProductId] = rowWarehouseStocks
		productsIDs = append(productsIDs, rowWarehouseStocks.ProductId)
	}

	var products = []domain.Product{}
	if err = tx.Set("gorm:query_option", "FOR UPDATE").Where("id IN (?)", productsIDs).Find(&products).Error; err != nil {
		return err
	}

	for _, product := range products {
		switch active {
		case true:
			product.AvailableQuantity += mapWarehouseStock[product.ID].Quantity
		case false:
			product.AvailableQuantity -= mapWarehouseStock[product.ID].Quantity
		}
		if err := tx.Save(&product).Error; err != nil {
			return pkgErrors.ProductUpdateFailedError(fmt.Sprintf(pkgErrors.ErrProductUpdateFailed, product.ID, err))
		}
	}

	return tx.Commit().Error

}

func (r ProductRepository) UpdateStockDeducted(ctx context.Context, orderProducts *[]domain.OrderProduct) (totalAmount float64, err error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	mapProductsID := map[string]domain.OrderProduct{}
	productsIDs := []string{}
	for _, orderProduct := range *orderProducts {
		mapProductsID[orderProduct.ProductID] = orderProduct
		productsIDs = append(productsIDs, orderProduct.ProductID)
	}

	var products = []domain.Product{}
	if err = tx.Set("gorm:query_option", "FOR UPDATE").Where("id IN (?)", productsIDs).Find(&products).Error; err != nil {
		return 0, err
	}

	for _, product := range products {
		if product.AvailableQuantity < mapProductsID[product.ID].Quantity {
			return 0, pkgErrors.OrderExceedStockError(fmt.Sprintf(pkgErrors.ErrOrderExceedStockWithIdAndCurrentStock, product.ID, product.AvailableQuantity))
		}

		totalAmount += product.Price
		product.AvailableQuantity -= mapProductsID[product.ID].Quantity
		product.ReservedQuantity += mapProductsID[product.ID].Quantity
		if err := tx.Save(&product).Error; err != nil {
			return 0, pkgErrors.ProductUpdateFailedError(fmt.Sprintf(pkgErrors.ErrProductUpdateFailed, product.ID, err))
		}
	}

	return totalAmount, tx.Commit().Error
}

func (r ProductRepository) UpdateStockRevert(ctx context.Context, orderItems *[]domain.OrderItem) (err error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	mapOrderItem := map[string]domain.OrderItem{}
	productsIDs := []string{}
	for _, orderItem := range *orderItems {
		mapOrderItem[orderItem.ProductId] = orderItem
		productsIDs = append(productsIDs, orderItem.ProductId)
	}

	var products = []domain.Product{}
	if err = tx.Set("gorm:query_option", "FOR UPDATE").Where("id IN (?)", productsIDs).Find(&products).Error; err != nil {
		return err
	}

	for _, product := range products {
		product.AvailableQuantity += mapOrderItem[product.ID].Quantity
		product.ReservedQuantity -= mapOrderItem[product.ID].Quantity
		if err := tx.Save(&product).Error; err != nil {
			return pkgErrors.ProductUpdateFailedError(fmt.Sprintf(pkgErrors.ErrProductUpdateFailed, product.ID, err))
		}
	}

	return tx.Commit().Error
}

func (r ProductRepository) FindByID(ctx context.Context, ID string) (data domain.Product, err error) {
	dbResult := r.DB.Model(&data).Where("id = ?", ID).Find(&data)
	if dbResult.RowsAffected == 0 {
		err = pkgErrors.ProductNotFoundError(fmt.Sprintf(pkgErrors.ErrProductNotFoundWithID, ID))
		return
	}

	return data, err
}

func (r ProductRepository) GetAll(ctx context.Context, params *domain.GetAllProductsPayload) (products []domain.Product, count int64, err error) {
	db := r.DB.WithContext(ctx).Limit(10).Offset(0)

	db, err = r.setQueryForGetAll(db, params)
	if err != nil {
		return
	}

	db = db.Where("available_quantity > 0")

	db = db.Find(&products).Offset(-1).Limit(-1).Count(&count) // reset offset & limit for the count
	if err = db.Error; err != nil {
		// https://www.postgresql.org/docs/current/errcodes-appendix.html
		postgresError, ok := db.Error.(*pgconn.PgError)
		if ok && postgresError.Code == "42703" {
			err = pkgErrors.InvalidColumnError(postgresError.Message)
		} else if err == gorm.ErrRecordNotFound {
			err = pkgErrors.ProductNotFoundError(pkgErrors.ErrProductNotFound)
		}
		return
	}

	return products, count, err
}

func (r ProductRepository) setQueryForGetAll(db *gorm.DB, params *domain.GetAllProductsPayload) (queryDB *gorm.DB, err error) {

	if params != nil {
		db = r.setLimitAndOffsetGetAll(db, params.Limit, params.Offset)
		db = r.setConditionForGetAll(db, "id", params.ID)
		db = r.setConditionForGetAll(db, "name", params.Name)
		db = r.setConditionForGetAll(db, "description", params.Description)

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

func (r ProductRepository) setLimitAndOffsetGetAll(db *gorm.DB, limit, offset int) *gorm.DB {
	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}
	return db
}

func (r ProductRepository) setConditionForGetAll(db *gorm.DB, field, criteria string) *gorm.DB {
	if criteria != "" {
		query := fmt.Sprintf("%s ILIKE '%s%s%s'", field, "%", criteria, "%")
		db = db.Where(query)
	}
	return db
}
