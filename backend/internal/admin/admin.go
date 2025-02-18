package admin

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// Load secret key from environment variable
var secretKey = []byte(os.Getenv("JWT_SECRET"))

func init() {
	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET is not set in environment variables")
	}
}

type Admin struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JWTResponse struct {
	Token string `json:"token"`
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
		"SELECT admin_id, full_name, email, password_hash, role, status FROM admins WHERE email = $1",
		admin.Email,
	).Scan(&dbAdmin.ID, &dbAdmin.FullName, &dbAdmin.Email, &hashedPassword, &dbAdmin.Role, &dbAdmin.Status)

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

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dbAdmin.Email,
		"role":  dbAdmin.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JWTResponse{Token: tokenString})
}

func NewPGAdminRepository(db *pgxpool.Pool) *PGAdminRepository {
	return &PGAdminRepository{db: db}
}
