CREATE TABLE Products (
        product_type_id VARCHAR(50),
        pid VARCHAR(50) PRIMARY KEY,
        poster VARCHAR(50),
        name VARCHAR(255) NOT NULL,
        sku VARCHAR(50) UNIQUE,
        price INT NOT NULL,
        quantity INT NOT NULL,
        brand VARCHAR(50),
        category VARCHAR(50),
        color VARCHAR(50),
        material VARCHAR(50),
        weight VARCHAR(50),
        size VARCHAR(50),
        original_price VARCHAR(50),
        sale BOOLEAN DEFAULT TRUE,
        discount INT DEFAULT 0,
        is_product BOOLEAN DEFAULT TRUE
);