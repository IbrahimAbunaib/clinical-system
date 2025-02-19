CREATE DATABASE payment_database;

USE payment_database;

CREATE TABLE IF NOT EXISTS payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    date DATETIME NOT NULL,
    notes TEXT,
    payment_method VARCHAR(50) NOT NULL
);
