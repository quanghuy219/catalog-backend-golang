package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string

	CategoryID int
	Category   Category

	UserID int
	User   User
}
