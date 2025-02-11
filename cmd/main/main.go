package main

import (
	"log"
	"lotus-task/internal/app/db"
	"github.com/gin-gonic/gin"
)


func init() {
	err := db.Connect()
	if err != nil {
		log.Println(err)
	}
	db.RunMigrations(db.DB)
	log.Println("successful to run migrations!")
}


func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PPPPPp",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
