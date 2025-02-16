package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Admin represents an admin user
type Admin struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// AdminRepository defines required methods for admin database operations
type AdminRepository interface {
	GetAdminByID(ctx context.Context, id int) (*Admin, error)
	CreateAdmin(ctx context.Context, admin Admin) error
}

// PGAdminRepository implements AdminRepository using PostgreSQL
type PGAdminRepository struct {
	db *pgxpool.Pool
}

// NewPGAdminRepository creates a new instance of PGAdminRepository
func NewPGAdminRepository(db *pgxpool.Pool) *PGAdminRepository {
	return &PGAdminRepository{db: db}
}

// GetAdminByID fetches an admin by ID
func (repo *PGAdminRepository) GetAdminByID(ctx context.Context, id int) (*Admin, error) {
	var admin Admin
	err := repo.db.QueryRow(ctx, "SELECT id, name, email FROM admin WHERE id=$1", id).Scan(&admin.ID, &admin.Name, &admin.Email)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// CreateAdmin inserts a new admin into the database
func (repo *PGAdminRepository) CreateAdmin(ctx context.Context, admin Admin) error {
	_, err := repo.db.Exec(ctx, "INSERT INTO admin (name, email) VALUES ($1, $2)", admin.Name, admin.Email)
	return err
}

// GetAdminHandler - HTTP handler for fetching an admin
func (repo *PGAdminRepository) GetAdminHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid admin ID", http.StatusBadRequest)
		return
	}

	admin, err := repo.GetAdminByID(req.Context(), id)
	if err != nil {
		http.Error(w, "Admin not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(admin)
}

// CreateAdminHandler - HTTP handler for adding a new admin
func (repo *PGAdminRepository) CreateAdminHandler(w http.ResponseWriter, req *http.Request) {
	var admin Admin
	err := json.NewDecoder(req.Body).Decode(&admin)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = repo.CreateAdmin(req.Context(), admin)
	if err != nil {
		http.Error(w, "Failed to create admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Admin created successfully")
}
