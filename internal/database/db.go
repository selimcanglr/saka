package database

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, ".book-cli.db")

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} 

	DB.AutoMigrate(
		&Book{}, 
		&BookRating{},
		&BookLog{},
	)
}