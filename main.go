package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"clinicalsystem/admin"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Global database connection pool
var pool *pgxpool.Pool

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get the database URL from environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in .env file or environment variables")
	}

	// Connect to PostgreSQL using pgxpool
	pool, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	// Check database connection
	var version string
	err = pool.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Failed to get PostgreSQL version: %v", err)
	}
	fmt.Println("✅ Connected to PostgreSQL, version:", version)

	// Initialize the router
	router := mux.NewRouter()

	// Home route
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// Create an instance of the admin repository
	adminRepo := admin.NewPGAdminRepository(pool)

	// Define admin routes
	router.HandleFunc("/admin/{id}", adminRepo.GetAdminHandler).Methods("GET")
	router.HandleFunc("/admin", adminRepo.CreateAdminHandler).Methods("POST")

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port
	}
	fmt.Println("🚀 Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// HomeHandler - Basic route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Clinical System API")
}
