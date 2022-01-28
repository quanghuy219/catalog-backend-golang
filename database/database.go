package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init database connection
func Init() *gorm.DB {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
