package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	Author    string `json:"author"`
	Synopsis  string `json:"synopsis"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
