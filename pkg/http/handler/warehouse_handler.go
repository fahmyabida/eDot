package handler

import (
	"net/http"

	"github.com/fahmyabida/eDot/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type WarehouseHandler struct {
	WarehouseUsecase domain.IWarehouseUsecase
}

func InitWarehouseHandler(e *echo.Group, WarehouseUsecase domain.IWarehouseUsecase) {
	handler := WarehouseHandler{WarehouseUsecase: WarehouseUsecase}

	e.POST("/warehouses/activation", handler.SwitchWarehouse)
	e.POST("/warehouses/stocks/transfer", handler.TransferStock)
}

func (h *WarehouseHandler) SwitchWarehouse(c echo.Context) error {

	var payload domain.SwitchWarehouseRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	response, err := h.WarehouseUsecase.SwitchWarehouse(ctx, &payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (h *WarehouseHandler) TransferStock(c echo.Context) error {

	var payload domain.StockTransferRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	response, err := h.WarehouseUsecase.TransferStockWarehouse(ctx, &payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
