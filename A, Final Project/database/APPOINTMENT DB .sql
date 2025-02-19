
CREATE DATABASE appointment;
USE appointment;
CREATE TABLE patients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL
);

CREATE TABLE doctors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    specialty VARCHAR(100) NOT NULL
);

CREATE TABLE appointments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    patient_id INT NOT NULL,
    doctor_id INT NOT NULL,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    appointment_status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_patient FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE,
    CONSTRAINT fk_doctor FOREIGN KEY (doctor_id) REFERENCES doctors(id) ON DELETE CASCADE
);

INSERT INTO patients (name, email, phone) VALUES 
('Ahmed Ali', 'ahmed@example.com', '01012345678'),
('Sara Mohamed', 'sara@example.com', '01098765432');

INSERT INTO doctors (name, specialty) VALUES 
('Dr. Mohamed Hassan', 'Cardiology'),
('Dr. Fatima Youssef', 'Dermatology');

INSERT INTO appointments (patient_id, doctor_id, appointment_date, appointment_time, appointment_status) VALUES 
(1, 1, '2025-02-20', '10:00:00', 'Confirmed'),
(2, 2, '2025-02-21', '14:30:00', 'Pending');
