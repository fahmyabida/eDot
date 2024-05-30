BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
	"id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('user-', uuid_generate_v4())),
	"username" VARCHAR(255) NOT NULL,
	"email" VARCHAR(255),
	"phone_number" VARCHAR(255),
	"password" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE UNIQUE INDEX "user_email" 
	ON users (email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX unique_phone_number 
	ON users (phone_number) WHERE phone_number IS NOT NULL;


CREATE TABLE IF NOT EXISTS products (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('product-', uuid_generate_v4())),
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "price" DECIMAL(10, 2) NOT NULL,
	"available_quantity" INT NOT NULL,
	"reserved_quantity" INT NOT NULL,
	"sold_quantity" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS orders (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('order-', uuid_generate_v4())),
    "user_id" VARCHAR(255) NOT NULL,
    "order_date" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    "total_amount" DECIMAL(10, 2) NOT NULL,
    "status" TEXT,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS order_items (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('order_item-', uuid_generate_v4())),
    "order_id" VARCHAR(255) NOT NULL,
    "product_id" VARCHAR(255) NOT NULL,
    "quantity" INT NOT NULL,
    "unit_price" DECIMAL(10, 2) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS warehouses (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('warehouse-', uuid_generate_v4())),
    "name" VARCHAR(255) NOT NULL,
    "location" VARCHAR(255),
    "status" VARCHAR(10) NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE')),
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS warehouse_stocks (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('warehouse_stock-', uuid_generate_v4())),
    "warehouse_id" VARCHAR(255) NOT NULL,
    "product_id" VARCHAR(255) NOT NULL,
    "quantity" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS shops (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('shop-', uuid_generate_v4())),
    "name" VARCHAR(255) NOT NULL,
    "location" VARCHAR(255),
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS shop_warehouses (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (concat('shop_warehouse-', uuid_generate_v4())),
    "shop_id" VARCHAR(255) NOT NULL,
    "warehouse_id" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
	"deleted_at" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    FOREIGN KEY (shop_id) REFERENCES shops(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

COMMIT;