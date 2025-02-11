package validators

import (
	"errors"
	"gorm.io/gorm"
)

func ValidateBlog(body string, db *gorm.DB) error {
	if len(body) < 3 || len(body) > 64 {
		return errors.New("blog's body must be between 3 and 64 characters")
	}
	return nil
}
