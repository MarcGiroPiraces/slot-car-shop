package db

import (
	"fmt"
	"go-slotcars/internal/cars"
	"go-slotcars/internal/users"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB initializes the SQLite connection
func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("slotcars.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Automatically migrate the schema (create tables)
	err = db.AutoMigrate(&users.User{}, &cars.Car{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate schema: %w", err)
	}

	log.Println("📦 Database connected and schema migrated successfully!")
	return db, nil
}
