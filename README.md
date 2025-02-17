for diagnostics you need to: 

1. Have Visual Studio installed, download the folder and open it on Visual Studio to test and run

2. Have localhost/phpMyAdmin, details are... servername: localhost; username = root; password: N/A; dbname: medical_diagnostics

3. When running the code by Visual Studio the login credentials for doctor are Username: doctor; Password: doctor123 

4. To login as patient you need to register first then log in back with the ID and password 

5. There are some data already stored in the database for patients check their ID if you want to log in as them the password is 123 

6. passwords can't be shown in the database because we set an encryption to the passwords in the database in the form of a privacy for patients' tests details

User Guide for the Doctor and Patient Modules:
This guide explains how to use the Doctor and Patient Modules in a simple way. Follow these steps to set up and test your system.



Using the Patient Module
1Register as a new patient by clicking "Register / Create Account" and filling in your details.
2 Log in using the email and password you provided.
3 Book an appointment with a doctor by selecting the "Book Appointment" button.
4 View medical history by clicking "Medical History", where you can see past diagnoses and prescriptions.
5 Receive notifications about upcoming appointments and updates from doctors.
 Using the Doctor Module
1 Login with the doctor’s credentials
2 View the daily schedule by clicking "View Daily Schedule" to check patient appointments.
3 Conduct online consultations by selecting "Conduct Online Consultations" and adding consultation details.
4 Request diagnostic tests for patients via "Request Diagnostic Tests".
5 Write and save prescriptions for patients under "Write & Save Prescriptions".  
I'll now create the Appointment Module, including:

Patient Booking Interface: Allows patients to select a doctor and schedule an appointment.
Doctor’s Appointment Management: Enables doctors to view and manage their appointments.
MySQL Database Integration: Stores and retrieves appointment details.

Navigating from the Main Page
1 When you open the main page, you can choose between Patient or Doctor Module.
2 Click on the respective button to go to the module you need.
3 Inside each module, you will see different options like Appointments, Diagnostics, and Prescriptions.


# 🏥 Clinical System - Admin Module

## 📌 Overview
This is the **Admin Module** for the Clinical System, built using **Golang**, **PostgreSQL**, and **REST API**. It provides authentication, user management, and admin functionalities.

---

## 🚀 Features
✅ Admin authentication (Login with JWT)  
✅ Secure password storage (Bcrypt hashing)  
✅ PostgreSQL database integration  
✅ RESTful API endpoints  

---

## 🛠️ Tech Stack
- **Backend:** Golang
- **Database:** PostgreSQL
- **Authentication:** JWT & Bcrypt
- **Frameworks/Libraries:** Gorilla Mux, Pgx, godotenv

---

