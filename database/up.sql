DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id VARCHAR(32) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS products;

CREATE TABLE products (
    id VARCHAR(32) PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(32) NOT NULL,
    FOREIGN key (created_by) REFERENCES users(id)
);