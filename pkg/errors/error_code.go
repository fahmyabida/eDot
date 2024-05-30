package errors

const (
	DuplicateUser              = "DUPLICATE_USER"
	UserNotFound               = "USER_NOT_FOUND"
	InvalidAmount              = "INVALID_AMOUNT"
	OrderNotFound              = "ORDER_NOT_FOUND"
	WarehouseStockNotFound     = "WAREHOUSE_STOCK_NOT_FOUND"
	WarehouseNotFound          = "WAREHOUSE_NOT_FOUND"
	OrderItemNotFound          = "ORDER_ITEM_NOT_FOUND"
	InvalidColumn              = "INVALID_COLUMN"
	ProductNotFound            = "PRODUCT_NOT_FOUND"
	OrderExceedStock           = "ORDER_EXCEEDED_STOCK"
	DuplicateOrder             = "DUPLICATE_ORDER"
	DuplicateOrderItem         = "DUPLICATE_ORDER_ITEM"
	ProductUpdateFailed        = "PRODUCT_UPDATE_FAILED"
	WarehouseStockUpdateFailed = "WAREHOUSE_STOCK_UPDATE_FAILED"
)
