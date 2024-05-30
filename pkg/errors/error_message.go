package errors

const (
	ErrUserNotFound                          = "User is not found with with email '%v' or phone number '%v'"
	ErrColumnInvalid                         = "Column is invalid"
	ErrDuplicateUser                         = "User is duplicated, either email or phone number"
	ErrDuplicateOrder                        = "Order is duplicated"
	ErrWarehouseStockNotFound                = "Warehouse Stock is not found"
	ErrWarehouseNotFoundWithID               = "Warehouse is not found with id '%v'"
	ErrDuplicateOrderItem                    = "Order Item is duplicated"
	ErrProductNotFoundWithID                 = "Product is not found with id '%v'"
	ErrProductNotFound                       = "Product is not found"
	ErrOrderNotFound                         = "Order is not found"
	ErrOrderItemNotFound                     = "Order item is not found"
	ErrOrderExceedStockWithIdAndCurrentStock = "Order exceed stock, product '%v' stock is '%v'"
	ErrProductUpdateFailed                   = "Error while update product with id '%v' & error '%v'"
	ErrWarehouseStockUpdateFailed            = "Error while update warehouse_stock with id '%v' & error '%v'"
	ErrInvalidQuantityWithID                 = "Exceed availiability quantity of product in warehouse stock with id '%v'"
)
