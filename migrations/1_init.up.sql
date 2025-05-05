CREATE TABLE products
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(64) NOT NULL,
    image_url  TEXT,
    cost       REAL        NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
CREATE TABLE carts
(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER NOT NULL,
    total_quantity INTEGER NOT NULL DEFAULT 0,
    created_at     TIMESTAMP,
    updated_at     TIMESTAMP,
    deleted_at     TIMESTAMP NULL
);
CREATE TABLE cart_items
(
    id         SERIAL PRIMARY KEY,
    cart_id    INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity   INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE CASCADE
);
CREATE TABLE orders
(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER      NOT NULL,
    cost           REAL         NOT NULL,
    raw_cost       REAL         NOT NULL,
    items_quantity INTEGER      NOT NULL,
    status         VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP,
    completed_at   TIMESTAMP NULL,
    updated_at     TIMESTAMP,
    deleted_at     TIMESTAMP NULL
);
CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity   INTEGER NOT NULL,
    cost       REAL    NOT NULL,
    raw_cost   REAL    NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);