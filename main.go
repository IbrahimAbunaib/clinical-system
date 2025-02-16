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
var db *pgxpool.Pool

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Debugging: Print env variables to verify
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	// Get the database URL from environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in .env file or environment variables")
	}

	// Connect to PostgreSQL
	db, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	// Check database connection
	var version string
	err = db.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Failed to get PostgreSQL version: %v", err)
	}
	fmt.Println("Connected to PostgreSQL, version:", version)

	// Initialize the router
	router := mux.NewRouter()

	// Home route
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// Create an instance of the admin repository
	adminRepo := admin.NewPGAdminRepository(db)

	// Define admin routes
	router.HandleFunc("/admin/{id}", adminRepo.GetAdminHandler).Methods("GET")
	router.HandleFunc("/admin", adminRepo.CreateAdminHandler).Methods("POST")

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// HomeHandler - Basic route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Clinical System API")
}
