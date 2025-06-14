CREATE TABLE User (
        user_id INT AUTO_INCREMENT PRIMARY KEY,
        user_name VARCHAR(100) NOT NULL,
        user_password VARCHAR(255) NOT NULL,
        age INT CHECK (age >= 0),
        address TEXT,
        user_email VARCHAR(100) NOT NULL UNIQUE
        
);
