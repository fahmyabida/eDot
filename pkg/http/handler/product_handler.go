package handler

import (
	"net/http"

	"github.com/fahmyabida/eDot/internal/app/domain"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUsecase domain.IProductUsecase
}

func InitProductHandler(e *echo.Group, ProductUsecase domain.IProductUsecase) {
	handler := ProductHandler{ProductUsecase: ProductUsecase}

	e.GET("/products", handler.GetAllProducts)
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {

	var payload domain.GetAllProductsPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	response, err := h.ProductUsecase.GetAllProducts(ctx, &payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
