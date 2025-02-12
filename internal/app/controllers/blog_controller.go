package controllers

import (
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/models"
	"lotus-task/internal/app/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	if validators.ValidateBlog(body.Body, db.DB) != nil {
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid body for blog",
			})
			return
		}
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

func UserLieksBlog(c *gin.Context) {
	var body struct {
		BlogID int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	user, _ := c.Get("user") // in middleware we already check if user exists!
	//check if blog exists
	var blog models.Blog
	foundBlog := db.DB.First(&blog, body.BlogID)
	if foundBlog.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog not found!",
		})
		return
	}
	//check if user hasn't already liked blog
	var userLikeBlog models.UserLikeBlog
	result := db.DB.Where("user_id = ? AND blog_id = ?", user.(models.User).ID, body.BlogID).First(&userLikeBlog)
	if result.Error == gorm.ErrRecordNotFound {
		db.DB.Create(&models.UserLikeBlog{
			UserID: user.(models.User).ID,
			BlogID: int(body.BlogID),
		})
		db.DB.Model(&models.Blog{}).Where("id = ?", body.BlogID).Update("likes_count", gorm.Expr("likes_count + ?", 1))
		c.JSON(http.StatusOK, gin.H{
			"message": "Blog liked successfully!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog already liked!",
	})

}
