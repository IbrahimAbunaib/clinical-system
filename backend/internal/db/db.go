package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// Load database credentials from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "MR.ibrahim2001"), // Should be from env for security
		getEnv("DB_NAME", "clinic"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	fmt.Println("✅ Connected to the database!")
}

// getEnv retrieves an environment variable or uses a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
