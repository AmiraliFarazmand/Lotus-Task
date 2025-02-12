package main

import (
	"log"
	"lotus-task/internal/app/controllers"
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/middleware"

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
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.ValidateIsAuthenticated)
	r.GET("/blogs", controllers.ListBlogs)
	r.GET("/blogs/:id", controllers.RetrieveBlog)
	r.POST("/blogs", middleware.RequireAuth, controllers.CreateBlog)
	r.POST("/like", middleware.RequireAuth, controllers.UserLieksBlog)
	r.Run() // listen and serve on 0.0.0.0:8080

}
