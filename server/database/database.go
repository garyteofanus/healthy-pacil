package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"healthy-pacil/config"
	"healthy-pacil/model"
	"log"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var dsn string
	if config.Config("DATABASE_URL") == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			config.Config("DB_HOST"),
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_NAME"),
			config.Config("DB_PORT"),
		)
	} else {
		dsn = config.Config("DATABASE_URL")
	}

	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db

	if dbError != nil {
		panic("failed to connect database")
	}

	migrateError := DB.AutoMigrate(
		&model.User{},
		&model.Appointment{})

	if migrateError != nil {
		panic("failed to auto migrate")
	}
}
