package usecase

import (
	"context"

	"github.com/fahmyabida/eDot/internal/app/domain"
)

type ProductUsecaseImpl struct {
	ProductRepo domain.IProductRepo
}

func NewProductUsecase(ProductRepo domain.IProductRepo) domain.IProductUsecase {
	return &ProductUsecaseImpl{
		ProductRepo: ProductRepo,
	}
}

func (u *ProductUsecaseImpl) GetAllProducts(ctx context.Context, payload *domain.GetAllProductsPayload) (domain.ProductGetAllResponse, error) {

	var response domain.ProductGetAllResponse

	data, total, err := u.ProductRepo.GetAll(ctx, payload)
	if err != nil {
		return response, err
	}

	response.Data = data
	response.Total = total

	return response, nil
}
