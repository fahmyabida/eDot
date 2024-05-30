package usecase

import (
	"context"
	"log"
	"strings"

	"github.com/fahmyabida/eDot/internal/app/domain"
)

type WarehouseUsecaseImpl struct {
	ProductRepo        domain.IProductRepo
	WarehouseRepo      domain.IWarehouseRepo
	WarehouseStockRepo domain.IWarehouseStockRepo
}

func NewWarehouseUsecase(ProductRepo domain.IProductRepo, WarehouseRepo domain.IWarehouseRepo, WarehouseStockRepo domain.IWarehouseStockRepo) domain.IWarehouseUsecase {
	return &WarehouseUsecaseImpl{
		ProductRepo:        ProductRepo,
		WarehouseRepo:      WarehouseRepo,
		WarehouseStockRepo: WarehouseStockRepo,
	}
}

func (u *WarehouseUsecaseImpl) SwitchWarehouse(ctx context.Context, payload *domain.SwitchWarehouseRequest) (response domain.SwitchWarehouseResponse, err error) {
	warehouse, err := u.WarehouseRepo.FindByID(ctx, payload.WarehouseID)
	if err != nil {
		log.Default().Printf("error occurred during get warehouse by id '%v': %v", payload.WarehouseID, err)
		return response, err
	}

	if warehouse.Status == string(payload.Mode) {
		response.Message = "currently warehouse is " + warehouse.Status
		response.Status = domain.WarehouseStatus(warehouse.Status)
		return response, nil
	}

	warehouse.Status = string(payload.Mode)
	err = u.WarehouseRepo.Update(ctx, &warehouse)
	if err != nil {
		log.Default().Printf("error occurred during update warehouse status by id '%v': %v", payload.WarehouseID, err)
		return response, err
	}

	warehouseStocks, _, err := u.WarehouseStockRepo.GetAll(ctx, &domain.GetAllWarehouseStocksPayload{
		WarehouseID: warehouse.ID,
	})
	if err != nil {
		log.Default().Printf("error occurred during get warehouse stock by id '%v': %v", warehouse.ID, err)
		return response, err
	}

	active := warehouse.Status == string(domain.ACTIVE)
	switch warehouse.Status {
	case string(domain.ACTIVE):
		err = u.ProductRepo.UpdateStockWarehouse(ctx, active, warehouseStocks)
	case string(domain.INACTIVE):
		err = u.ProductRepo.UpdateStockWarehouse(ctx, active, warehouseStocks)
	}
	if err != nil {
		log.Default().Printf("error occurred during update product quantity while warehouse is switch to '%v': %v", warehouse.Status, err)
		return response, err
	}

	response.Message = "warehouse switched into " + warehouse.Status
	response.Status = domain.WarehouseStatus(warehouse.Status)
	return
}

func (u *WarehouseUsecaseImpl) TransferStockWarehouse(ctx context.Context, payload *domain.StockTransferRequest) (response domain.StockTransferResponse, err error) {
	productIDs := []string{}
	for _, row := range payload.Products {
		productIDs = append(productIDs, row.ProductID)
	}

	warehouseStocks, err := u.WarehouseStockRepo.GetByWarehouseIAndProductIds(ctx, payload.SourceWarehouseID, productIDs)
	if err != nil {
		log.Default().Printf("error occurred during get warehouse stock with warehouse_id '%v' & product_ids '%v': %v",
			payload.SourceWarehouseID, strings.Join(productIDs, ", "), err)
		return response, err
	}

	err = u.WarehouseStockRepo.TransferStock(ctx, domain.SOURCE, warehouseStocks)
	if err != nil {
		log.Default().Printf("error occurred during update source warehouse while transfer stock with warehouse_id '%v' & product_ids '%v': %v",
			payload.SourceWarehouseID, strings.Join(productIDs, ", "), err)
		return response, err
	}

	transferStock := []domain.WarehouseStock{}
	for _, row := range warehouseStocks {
		transferStock = append(transferStock, domain.WarehouseStock{
			WarehouseId: row.WarehouseId,
			ProductId:   row.ProductId,
			Quantity:    row.Quantity,
		})
	}

	err = u.WarehouseStockRepo.TransferStock(ctx, domain.DESTINATION, transferStock)
	if err != nil {
		log.Default().Printf("error occurred during update destination warehouse while transfer stock with warehouse_id '%v' & product_ids '%v': %v",
			payload.SourceWarehouseID, strings.Join(productIDs, ", "), err)
		return response, err
	}

	response.Message = "success transfer stock"
	response.Status = "SUCCEEDED"
	return response, nil
}
