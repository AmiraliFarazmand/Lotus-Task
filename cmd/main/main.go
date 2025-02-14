package main

import (
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/routes"
)

func main() {
	db.InitDB()
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
