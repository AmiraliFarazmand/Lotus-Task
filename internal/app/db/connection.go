package db

import (
    "errors"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
    err := godotenv.Load(".env")
    if err != nil {
        return nil, errors.New(">ERR db.Connect(). Error loading .env file")
    }

    dsn := os.Getenv("DSN")
    if dsn == "" {
        return nil, errors.New(">ERR db.Connect(). DSN not found in environment variables")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, errors.New(">ERR db.Connect(). Failed to connect to database")
    }
    return db, nil
}