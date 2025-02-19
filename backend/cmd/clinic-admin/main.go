package main

import (
	"log"
	"os"

	"github.com/IbrahimAbunaib/clinical-system/backend/db"
	"github.com/IbrahimAbunaib/clinical-system/backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}

	// Connect to database
	db.ConnectDB()

	router := gin.Default()

	// Register routes
	routes.AdminRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("🚀 Server is running on port " + port)
	router.Run(":" + port)
}
