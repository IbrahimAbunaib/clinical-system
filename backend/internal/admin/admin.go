package admin

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET")) // From .env

type Admin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTResponse struct {
	Token string `json:"token"` // Match frontend expectation
}

type PGAdminRepository struct {
	db *pgxpool.Pool
}

func (repo *PGAdminRepository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var dbAdmin Admin
	var hashedPassword string
	err := repo.db.QueryRow(
		context.Background(),
		"SELECT id, email, password_hash FROM admins WHERE email = $1", // Updated query
		admin.Email,
	).Scan(&dbAdmin.ID, &dbAdmin.Email, &hashedPassword)
	if err != nil {
		log.Printf("Admin not found: %s (%v)", admin.Email, err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(admin.Password)); err != nil {
		log.Printf("Password mismatch for %s: %v", admin.Email, err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dbAdmin.Email,
		"admin": true,
	})

	tokenString, _ := token.SignedString(secretKey)
	json.NewEncoder(w).Encode(JWTResponse{Token: tokenString}) // Ensure "token" key
}

func NewPGAdminRepository(db *pgxpool.Pool) *PGAdminRepository {
	return &PGAdminRepository{db: db}
}
