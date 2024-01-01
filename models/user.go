package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	Username  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Photos    []Photo   `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Metode untuk membuat dan memanipulasi model User akan ditambahkan di sini
