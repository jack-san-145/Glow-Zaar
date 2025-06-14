CREATE TABLE Orders (
        order_id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        status ENUM('cart', 'ordered') DEFAULT 'cart',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        ordered_at TIMESTAMP NULL
) AUTO_INCREMENT=100;