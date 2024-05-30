package handler

import (
	"net/http"

	"github.com/fahmyabida/eDot/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderUsecase domain.IOrderUsecase
}

func InitOrderHandler(e *echo.Group, OrderUsecase domain.IOrderUsecase) {
	handler := OrderHandler{OrderUsecase: OrderUsecase}

	e.POST("/orders", handler.CreateOrder)
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {

	var payload domain.CreateOrderRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	response, err := h.OrderUsecase.Create(ctx, &payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
