package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/minhtridinh/trid-profile-go/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	// Extract database name from DSN
	dsnParts := strings.Split(dsn, "/")
	dbName := strings.Split(dsnParts[len(dsnParts)-1], "?")[0]

	// Create a DSN without database name for initial connection
	rootDSN := strings.Join(dsnParts[:len(dsnParts)-1], "/") + "/"

	// Connect to MySQL without specifying database
	sqlDB, err := sql.Open("mysql", rootDSN)
	if err != nil {
		log.Fatal("Failed to connect to MySQL server:", err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)

	// Create database if it doesn't exist
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	fmt.Printf("Database '%s' ensured to exist\n", dbName)

	// Now connect with the database specified
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&model.User{}, &model.Profile{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	return db
}
