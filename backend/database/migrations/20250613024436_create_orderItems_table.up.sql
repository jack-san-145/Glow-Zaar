CREATE TABLE OrderItems (
        order_id INT ,
        pid VARCHAR(50) NOT NULL,
        quantity INT NOT NULL,
        price INT NOT NULL,
        added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);