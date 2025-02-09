package main

import (
	"fmt"
	"log"
	"lotus-task/internal/app/db"
)

func main() {
	fmt.Println()
	_, err := db.Connect()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connected to database")
}
