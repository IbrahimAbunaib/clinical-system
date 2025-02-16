package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your-secret-key")

type Admin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTResponse struct {
	Token string `json:"token"`
}

type PGAdminRepository struct {
	db *pgxpool.Pool
}

// LoginHandler handles admin login
func (repo *PGAdminRepository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var dbAdmin Admin
	var hashedPassword string
	err = repo.db.QueryRow(context.Background(), "SELECT id, email, password FROM admin WHERE email=$1", admin.Email).
		Scan(&dbAdmin.ID, &dbAdmin.Email, &hashedPassword)
	if err != nil {
		fmt.Println("❌ Admin not found:", admin.Email)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Debugging output
	fmt.Println("✅ Admin found in DB:", dbAdmin.Email)
	fmt.Println("🔹 Stored Hash:", hashedPassword)
	fmt.Println("🔹 Entered Password:", admin.Password)

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(admin.Password)); err != nil {
		fmt.Println("❌ Password comparison failed:", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dbAdmin.Email,
		"admin": true,
	})

	tokenString, _ := token.SignedString(secretKey)

	json.NewEncoder(w).Encode(JWTResponse{Token: tokenString})
}

// NewPGAdminRepository initializes a new admin repository
func NewPGAdminRepository(db *pgxpool.Pool) *PGAdminRepository {
	return &PGAdminRepository{db: db}
}
