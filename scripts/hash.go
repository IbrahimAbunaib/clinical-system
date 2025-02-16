package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "MR.ibrahim2001" // Plain text password

	// Generate a hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("🔐 Hashed password:", string(hashedPassword))
}
