package main

import (
	"fmt"
	"github.com/minhtridinh/trid-profile-go/internal/config"
	"github.com/minhtridinh/trid-profile-go/internal/database"
	"log"
)

func main() {
	// Load configuration from .env file
	cfg := config.LoadConfig()

	// Initialize database connection and run migrations
	fmt.Println("Connecting to database and running migrations...")
	db := database.InitDB(cfg.DatabaseDSN)

	// Check database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Database connected successfully!")
	fmt.Println("Migrations completed successfully!")
}
