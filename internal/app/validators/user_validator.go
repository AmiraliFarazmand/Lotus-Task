package validators

import (
	"errors"
	"fmt"
	"lotus-task/internal/app/models"
	"gorm.io/gorm"
)

func checkUniquenessUsername(db *gorm.DB, username string) error {
	var user models.User
	if db.Where("username = ?", username).First(&user).Error == nil {
		return fmt.Errorf("username %s already exists", username)
	}
	return nil
}
func ValidateUsername(username string, db *gorm.DB) error {
	if len(username) < 3 || len(username) > 64 {
		return errors.New("username must be between 3 and 64 characters")
	}
	if err := checkUniquenessUsername(db, username); err != nil {
		return err
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if len(password) > 64 {
		return errors.New("password must be less than 64 characters long")
	}
	return nil
}
