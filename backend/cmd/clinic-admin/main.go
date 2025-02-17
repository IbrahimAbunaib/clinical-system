package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IbrahimAbunaib/clinical-system/backend/internal/admin"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Debugging: Print env variables
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	// Get the database URL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in .env file or environment variables")
	}

	// Connect to PostgreSQL
	dbPool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Check database connection
	var version string
	err = dbPool.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Failed to get PostgreSQL version: %v", err)
	}
	fmt.Println("✅ Connected to PostgreSQL, version:", version)

	// Initialize the router
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// Serve frontend files
	router.PathPrefix("/admin/").Handler(
		http.StripPrefix("/admin/", http.FileServer(http.Dir("../../frontend/admin"))),
	)

	// Initialize the admin repository
	adminRepo := admin.NewPGAdminRepository(dbPool)

	// Define admin routes
	router.HandleFunc("/admin/login", adminRepo.LoginHandler).Methods("POST")

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("🚀 Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// HomeHandler - Basic route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Clinical System API")
}
