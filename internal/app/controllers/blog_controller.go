package controllers

import (
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBlogs(c *gin.Context) {
	var blogs []models.Blog
	db.DB.Find(&blogs)
	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func RetrieveBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	result := db.DB.First(&blog, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, blog)
}
func CreateBlog(c *gin.Context) {
	var body struct {
		Body string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	user, _ := c.Get("user") // in middleware we already check if user exists!

	db.DB.Create(&models.Blog{
		Body:   body.Body,
		UserID: user.(models.User).ID,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog created!",
	})
}
