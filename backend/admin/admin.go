package models

import (
	"backend/db"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Hash the password before saving to DB
func (a *Admin) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}

// Authenticate an admin by verifying password
func AuthenticateAdmin(email, password string) (*Admin, error) {
	var admin Admin
	err := db.DB.QueryRow("SELECT id, email, password, role FROM admins WHERE email=$1", email).
		Scan(&admin.ID, &admin.Email, &admin.Password, &admin.Role)
	if err != nil {
		return nil, err
	}

	// Check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &admin, nil
}
