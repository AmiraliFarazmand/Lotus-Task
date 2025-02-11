package db

import (
	"errors"
	"fmt"
	"lotus-task/internal/app/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	dsn, err := utils.ReadEnv("DSN")
	if err != nil {
		return nil, fmt.Errorf(">ERR db.Connect().%w", err)
	}
    
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(">ERR db.Connect(). Failed to connect to database")
	}

	// Assign the global variable DB |   TODO: can be changed in the future
	DB = db
	return db, nil
}
