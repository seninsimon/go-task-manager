package db

import (
	"log"

	"task-manager/internal/config"
	"task-manager/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	database, err := gorm.Open(postgres.Open(cfg.DatabaseDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(
		&models.User{},
	    &models.Task{},
	)
	if err == nil {
		log.Println("Database migrated successfully")
	}
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	

	DB = database
	log.Println("Database connected successfully")
}
   