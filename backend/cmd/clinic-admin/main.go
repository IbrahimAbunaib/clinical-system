package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IbrahimAbunaib/clinical-system/backend/internal/admin"
	"github.com/IbrahimAbunaib/clinical-system/backend/internal/db"
	"github.com/IbrahimAbunaib/clinical-system/backend/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// init() runs before main() to load environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found. Using system environment variables.")
	}
}

func main() {
	// Ensure JWT_SECRET is set
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("❌ JWT_SECRET is not set in environment variables")
	}
	fmt.Println("✅ JWT_SECRET is set successfully!")

	// Connect to the database
	db.ConnectDB()

	// Load DATABASE_URL from environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set in .env file or environment variables")
	}

	// Connect to PostgreSQL
	dbPool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("❌ Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Verify database connection
	var version string
	err = dbPool.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("❌ Failed to get PostgreSQL version: %v", err)
	}
	fmt.Println("✅ Connected to PostgreSQL, version:", version)

	// Initialize Router
	router := mux.NewRouter()

	// Admin Routes
	apiRouter := router.PathPrefix("/api/admin").Subrouter()
	adminRepo := admin.NewPGAdminRepository(dbPool)

	// Admin Login Route (No Middleware)
	apiRouter.HandleFunc("/login", adminRepo.LoginHandler).Methods("POST")

	// ✅ Protect admin routes with JWT middleware
	protectedRoutes := apiRouter.PathPrefix("/protected").Subrouter()
	protectedRoutes.Use(middleware.JWTMiddleware)
	protectedRoutes.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Welcome to the Admin Dashboard!")
	}).Methods("GET")

	// Debugging: Print all registered routes
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			methods, _ := route.GetMethods()
			fmt.Println("🛤️ Registered Route:", methods, path)
		}
		return nil
	})

	// Start HTTP Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("🚀 Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
