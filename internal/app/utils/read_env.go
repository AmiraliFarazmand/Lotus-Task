package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv(lookupStr string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("error occured on loading .env file")
	}

	found := os.Getenv(lookupStr)
	if found == "" {
		return  "", errors.New(lookupStr + "not found in environment variables")
	}
	return found, nil
}
