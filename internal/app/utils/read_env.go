package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error occured on loading .env file")
	}
}


func ReadEnv(lookupStr string) (string, error) {

	found := os.Getenv(lookupStr)
	if found == "" {
		return "", errors.New(lookupStr + "not found in environment variables")
	}
	return found, nil
}
