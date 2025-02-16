-- Create database
CREATE DATABASE IF NOT EXISTS medical_diagnostics;
USE medical_diagnostics;

-- Create tables
CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS doctors (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS test_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    patient_id VARCHAR(50),
    doctor_id VARCHAR(50),
    test_type VARCHAR(50) NOT NULL,
    diagnostics_id VARCHAR(100),
    result TEXT NOT NULL,
    recommended_action TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    FOREIGN KEY (doctor_id) REFERENCES doctors(id)
);

-- Insert default doctor account
-- Password is 'doctor123' hashed
INSERT INTO doctors (id, username, password) 
VALUES ('DOC001', 'doctor', '$2y$10$8tdsR.2X0RxQeN6VsX/y0.AY.KZgRWxLEGYF.qhkY7TqDGTcLtvmq')
ON DUPLICATE KEY UPDATE username = username;
