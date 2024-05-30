package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fahmyabida/eDot/internal/app/domain"
	pkgErrors "github.com/fahmyabida/eDot/pkg/errors"
)

type OrderUsecaseImpl struct {
	ProductRepo   domain.IProductRepo
	OrderRepo     domain.IOrderRepo
	OrderItemRepo domain.IOrderItemRepo
}

func NewOrderUsecase(ProductRepo domain.IProductRepo, OrderRepo domain.IOrderRepo, OrderItemRepo domain.IOrderItemRepo) domain.IOrderUsecase {
	return &OrderUsecaseImpl{
		ProductRepo:   ProductRepo,
		OrderRepo:     OrderRepo,
		OrderItemRepo: OrderItemRepo,
	}
}

func (u *OrderUsecaseImpl) Create(ctx context.Context, request *domain.CreateOrderRequest) (response domain.CreteOrderResponse, err error) {
	for i, product := range request.Products {
		productData, err := u.ProductRepo.FindByID(ctx, product.ProductID)
		if err != nil {
			return response, pkgErrors.ProductNotFoundError(fmt.Sprintf(pkgErrors.ErrProductNotFoundWithID, product.ProductID))
		}
		request.Products[i].Price = productData.Price
	}

	totalAmount, err := u.ProductRepo.UpdateStockDeducted(ctx, &request.Products)
	if err != nil {
		log.Default().Printf("error occurred during stock update, deducted: %v", err)
		return
	}

	status := string(domain.ORDERED)
	order := domain.Order{
		UserID:      request.UserID,
		TotalAmount: totalAmount,
		Status:      status,
	}
	err = u.OrderRepo.Create(ctx, &order)
	if err != nil {
		log.Default().Printf("error occurred during create order: %v", err)
		return
	}

	orderItems := []domain.OrderItem{}
	for _, product := range request.Products {
		orderItems = append(orderItems, domain.OrderItem{
			OrderId:   order.ID,
			ProductId: product.ProductID,
			Quantity:  product.Quantity,
			UnitPrice: product.Price,
		})
	}

	err = u.OrderItemRepo.Create(ctx, &orderItems)
	if err != nil {
		log.Default().Printf("error occurred during create order_items: %v", err)
		return
	}

	response.Message = "products has been successfully ordered, please continue the payment"
	response.Status = status

	go u.CancelOrderBackground(context.Background(), order)

	return response, nil
}

func (u *OrderUsecaseImpl) CancelOrderBackground(ctx context.Context, order domain.Order) (err error) {
	log.Default().Printf("background process to cancel orderd with id '%v' is started...", order.ID)
	time.Sleep(1 * time.Minute)
	order, err = u.OrderRepo.FindByID(ctx, order.ID)
	if err != nil {
		log.Default().Printf("background process to cancel orderd with id '%v', error during get by id: %v", order.ID, err)
	} else if order.Status == string(domain.ORDERED) {
		err = u.cancelOrder(ctx, order)
		if err != nil {
			log.Default().Printf("background process to cancel orderd with id '%v', error: %v", order.ID, err)
		}
		log.Default().Printf("background process to cancel orderd with id '%v' is succeeded", order.ID)
	}
	log.Default().Printf("background process to cancel orderd with id '%v' is stopped...", order.ID)
	return
}

func (u *OrderUsecaseImpl) CancelExceededOrder(ctx context.Context) (err error) {
	orders, err := u.OrderRepo.GetOrdersByStatusAndCreatedAtLessThan(ctx, string(domain.ORDERED), time.Now().Add(-1*time.Minute))
	if err != nil {
		if err == pkgErrors.OrderNotFoundError(pkgErrors.ErrOrderNotFound) {
			log.Default().Println("no order exceeded")
			return nil
		}
		log.Default().Printf("error occurred during get order which exceeded order time: %v", err)
		return err
	}

	counter := 0
	for _, order := range orders {
		err = u.cancelOrder(ctx, order)
		if err != nil {
			log.Default().Printf("error occurred during update orders with id '%v': %v", order.ID, err)
			continue
		}
		log.Default().Printf("succeeded cancel exceeded order with id '%v'", order.ID)
		counter++
	}
	log.Default().Printf("succeeded cancel '%v' exceeded order ", counter)

	return nil
}

func (u *OrderUsecaseImpl) cancelOrder(ctx context.Context, order domain.Order) (err error) {
	order.Status = string(domain.TIMEOUT)
	err = u.OrderRepo.Update(ctx, &order)
	if err != nil {
		log.Default().Printf("error occurred during update orders: %v", err)
		return err
	}

	ordersItems, err := u.OrderItemRepo.GetOrderItemsByOrderId(ctx, order.ID)
	if err != nil {
		if err == pkgErrors.OrderItemNotFoundError(pkgErrors.ErrOrderItemNotFound) {
			return nil
		}
		log.Default().Printf("error occurred during update order items by order_id with order_id: '%v' & error: '%v'", order.ID, err)
		return err
	}

	err = u.ProductRepo.UpdateStockRevert(ctx, &ordersItems)
	if err != nil {
		log.Default().Printf("error occurred during revert stock in product data: %v", err)
		return err
	}
	return nil
}
