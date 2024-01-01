package models

import (
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	ID        uint      `gorm:"primary_key"`
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}


