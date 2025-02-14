package db

import (
	"log"
)

func InitDB() {
	err := connect()
	if err != nil {
		log.Println(err)
	}
	runMigrations(DB)
	log.Println("Database migrations completed!")
}
