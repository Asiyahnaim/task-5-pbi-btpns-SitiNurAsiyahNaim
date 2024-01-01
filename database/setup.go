package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"main/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open("postgre:12345(localhost:5432)"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
