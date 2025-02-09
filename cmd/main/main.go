package main

import (
	"fmt"
	"log"
	"lotus-task/internal/app/db"
)

func main() {
	fmt.Println()
	database, err := db.Connect()
	if err != nil {
		log.Println(err)
	}
	
	db.RunMigrations(database)
	fmt.Println("successful to run migrations!")
}
