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
    total_cost     FLOAT            DEFAULT 0,
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
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
);

-- === Products ===
INSERT INTO products (name, image_url, cost, created_at, updated_at)
VALUES ('Product A', 'https://example.com/a.jpg', 10.5, NOW(), NOW()),
       ('Product B', 'https://example.com/b.jpg', 15.0, NOW(), NOW()),
       ('Product C', 'https://example.com/c.jpg', 7.25, NOW(), NOW()),
       ('Product D', 'https://example.com/d.jpg', 21.99, NOW(), NOW()),
       ('Product E', 'https://example.com/e.jpg', 13.75, NOW(), NOW()),
       ('Product F', 'https://example.com/f.jpg', 9.5, NOW(), NOW()),
       ('Product G', 'https://example.com/g.jpg', 11.0, NOW(), NOW()),
       ('Product H', 'https://example.com/h.jpg', 17.3, NOW(), NOW()),
       ('Product I', 'https://example.com/i.jpg', 5.9, NOW(), NOW()),
       ('Product J', 'https://example.com/j.jpg', 12.6, NOW(), NOW());

-- === Carts & Cart Items & Orders & Order Items ===

-- USER 1
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (1, 6, 79.85, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (1, 1, 2, NOW(), NOW()),
       (1, 4, 1, NOW(), NOW()),
       (1, 5, 2, NOW(), NOW()),
       (1, 9, 1, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (1, 79.85, 6, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 1, 2, 10.5, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 4, 1, 21.99, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 5, 2, 13.75, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 9, 1, 5.9, NOW(), NOW());

-- USER 2
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (2, 7, 96.95, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (2, 2, 2, NOW(), NOW()),
       (2, 3, 1, NOW(), NOW()),
       (2, 6, 3, NOW(), NOW()),
       (2, 10, 1, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (2, 96.95, 7, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 2, 2, 15.0, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 3, 1, 7.25, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 6, 3, 9.5, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 10, 1, 12.6, NOW(), NOW());

-- USER 3
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (3, 5, 61.85, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (3, 2, 2, NOW(), NOW()),
       (3, 7, 2, NOW(), NOW()),
       (3, 9, 1, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (3, 61.85, 5, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 2, 2, 15.0, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 7, 2, 11.0, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 9, 1, 5.9, NOW(), NOW());

-- USER 4
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (4, 6, 91.05, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (4, 1, 1, NOW(), NOW()),
       (4, 4, 2, NOW(), NOW()),
       (4, 6, 3, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (4, 91.05, 6, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 1, 1, 10.5, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 4, 2, 21.99, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 6, 3, 9.5, NOW(), NOW());

-- USER 5
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (5, 7, 82.3, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (5, 3, 1, NOW(), NOW()),
       (5, 8, 3, NOW(), NOW()),
       (5, 5, 3, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (5, 82.3, 7, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 3, 1, 7.25, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 8, 3, 17.3, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 5, 3, 13.75, NOW(), NOW());

-- USER 6
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (6, 6, 69.9, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (6, 10, 3, NOW(), NOW()),
       (6, 7, 2, NOW(), NOW()),
       (6, 9, 1, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (6, 69.9, 6, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 10, 3, 12.6, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 7, 2, 11.0, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 9, 1, 5.9, NOW(), NOW());

-- USER 7
INSERT INTO carts (user_id, total_quantity, total_cost, created_at, updated_at)
VALUES (7, 5, 57.5, NOW(), NOW());
INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
VALUES (7, 1, 1, NOW(), NOW()),
       (7, 2, 2, NOW(), NOW()),
       (7, 6, 2, NOW(), NOW());
INSERT INTO orders (user_id, cost, items_quantity, status, created_at, updated_at, completed_at)
VALUES (7, 57.5, 5, 'completed', NOW(), NOW(), NOW());
INSERT INTO order_items (order_id, product_id, quantity, cost, created_at, updated_at)
VALUES ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 1, 1, 10.5, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 2, 2, 15.0, NOW(), NOW()),
       ((SELECT currval(pg_get_serial_sequence('orders', 'id'))), 6, 2, 9.5, NOW(), NOW());
