package database

import (
	"github.com/dinhminhtri/triD-profile/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	db.AutoMigrate(&model.User{}, &model.Profile{})
	return db
}
