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

-- Products
INSERT INTO products (name, image_url, cost, created_at)
VALUES
    ('Apple iPhone 14', 'https://example.com/iphone14.jpg', 799.99, NOW()),
    ('Samsung Galaxy S23', 'https://example.com/galaxys23.jpg', 699.99, NOW()),
    ('Google Pixel 8', 'https://example.com/pixel8.jpg', 599.99, NOW()),
    ('Sony WH-1000XM5', 'https://example.com/sonyheadphones.jpg', 349.99, NOW()),
    ('Apple MacBook Pro', 'https://example.com/macbookpro.jpg', 1999.99, NOW()),
    ('Dell XPS 13', 'https://example.com/dellxps.jpg', 1299.99, NOW()),
    ('Logitech MX Master 3', 'https://example.com/mouse.jpg', 99.99, NOW()),
    ('Apple Watch Series 9', 'https://example.com/applewatch.jpg', 399.99, NOW()),
    ('Amazon Echo Dot', 'https://example.com/echodot.jpg', 49.99, NOW()),
    ('Samsung QLED TV', 'https://example.com/samsungtv.jpg', 899.99, NOW());

-- Carts
INSERT INTO carts (user_id, total_quantity, created_at)
VALUES
    (1, 2, NOW()),
    (2, 3, NOW()),
    (3, 1, NOW()),
    (4, 4, NOW()),
    (5, 2, NOW()),
    (6, 3, NOW()),
    (7, 1, NOW()),
    (8, 5, NOW()),
    (9, 2, NOW()),
    (10, 3, NOW());

-- Cart items
INSERT INTO cart_items (cart_id, product_id, quantity, created_at)
VALUES
    (1, 1, 1, NOW()),
    (1, 2, 1, NOW()),
    (2, 3, 2, NOW()),
    (2, 4, 1, NOW()),
    (3, 5, 1, NOW()),
    (4, 6, 2, NOW()),
    (4, 7, 2, NOW()),
    (5, 8, 2, NOW()),
    (6, 9, 1, NOW()),
    (6, 10, 2, NOW()),
    (7, 1, 1, NOW()),
    (8, 2, 3, NOW()),
    (8, 3, 2, NOW()),
    (9, 4, 2, NOW()),
    (10, 5, 3, NOW());

-- Orders
INSERT INTO orders (user_id, cost, raw_cost, items_quantity, status, created_at, updated_at)
VALUES
    (1, 1599.98, 1599.98, 2, 'created', NOW(), NOW()),
    (2, 949.98, 949.98, 3, 'completed', NOW(), NOW()),
    (3, 1999.99, 1999.99, 1, 'created', NOW(), NOW()),
    (4, 2799.97, 2799.97, 4, 'shipped', NOW(), NOW()),
    (5, 199.98, 199.98, 2, 'cancelled', NOW(), NOW()),
    (6, 1449.98, 1449.98, 3, 'created', NOW(), NOW()),
    (7, 899.99, 899.99, 1, 'completed', NOW(), NOW()),
    (8, 1049.97, 1049.97, 5, 'processing', NOW(), NOW()),
    (9, 699.98, 699.98, 2, 'created', NOW(), NOW()),
    (10, 3899.97, 3899.97, 3, 'completed', NOW(), NOW());

-- Order items
INSERT INTO order_items (order_id, product_id, quantity, cost, raw_cost, created_at)
VALUES
    (1, 1, 1, 799.99, 799.99, NOW()),
    (1, 2, 1, 799.99, 799.99, NOW()),
    (2, 3, 2, 599.99, 599.99, NOW()),
    (2, 4, 1, 349.99, 349.99, NOW()),
    (3, 5, 1, 1999.99, 1999.99, NOW()),
    (4, 6, 2, 1299.99, 1299.99, NOW()),
    (4, 7, 2, 99.99, 99.99, NOW()),
    (5, 8, 2, 99.99, 99.99, NOW()),
    (6, 9, 1, 49.99, 49.99, NOW()),
    (6, 10, 2, 699.99, 699.99, NOW()),
    (7, 10, 1, 899.99, 899.99, NOW()),
    (8, 2, 3, 699.99, 699.99, NOW()),
    (8, 3, 2, 599.99, 599.99, NOW()),
    (9, 4, 2, 349.99, 349.99, NOW()),
    (10, 5, 3, 1299.99, 1299.99, NOW());
