package main

import (
	"fmt"
	"github.com/minhtridinh/trid-profile-go/internal/api"
	"github.com/minhtridinh/trid-profile-go/internal/config"
	"github.com/minhtridinh/trid-profile-go/internal/database"
	"github.com/minhtridinh/trid-profile-go/internal/repository"
	"github.com/minhtridinh/trid-profile-go/internal/service"
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

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize router
	router := api.SetupRouter(userService)

	// Start HTTP server
	port := cfg.Port
	if port == "" {
		port = "8080" // Default port if not specified in config
	}

	fmt.Printf("Starting server on port %s...\n", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
