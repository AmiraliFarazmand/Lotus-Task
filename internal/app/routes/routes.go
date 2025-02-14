package routes

import (
	"lotus-task/internal/app/controllers"
	"lotus-task/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.ValidateIsAuthenticated)
	r.GET("/blogs", controllers.ListBlogs)
	r.GET("/blogs/:id", controllers.RetrieveBlog)
	r.POST("/blogs", middleware.RequireAuth, controllers.CreateBlog)
	r.POST("/like", middleware.RequireAuth, controllers.UserLieksBlog)

	return r
}
