package db

import (
	"errors"
	"fmt"
	"lotus-task/internal/app/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn, err := utils.ReadEnv("DSN")
	if err != nil {
		return fmt.Errorf(">ERR db.Connect().%w", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(">ERR db.Connect(). Failed to connect to database")
	}

	// Assign the global variable DB
	DB = db
	return nil
}
