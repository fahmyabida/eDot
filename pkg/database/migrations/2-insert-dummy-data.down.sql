BEGIN;

DELETE FROM order_items WHERE 1 = 1;

DELETE FROM warehouse_stocks WHERE 1 = 1;

DELETE FROM warehouses WHERE 1 = 1;

DELETE FROM shop_warehouses WHERE 1 = 1;

DELETE FROM shops WHERE 1 = 1;

DELETE FROM orders WHERE 1 = 1;

DELETE FROM users WHERE 1 = 1;

COMMIT;