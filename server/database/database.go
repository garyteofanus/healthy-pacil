package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"healthy-pacil/models"
)

func Connect() {
	dsn := "host=localhost user=gorm password=gorm dbname=healthy-pacil port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbError != nil {
		panic("failed to connect database")
	}

	migrateError := db.AutoMigrate(
		&models.User{},
		&models.Administrator{},
		&models.Patient{},
		&models.Doctor{})

	if migrateError != nil {
		panic("failed to auto migrate")
	}
}
