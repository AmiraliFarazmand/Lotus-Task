package controllers

import (
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/models"
	"lotus-task/internal/app/utils"
	"lotus-task/internal/app/validators"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}
	if err := validators.ValidateUsernamePassword(body.Username, body.Password, db.DB); err != nil { //  validate username and password by our own logic
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash Password ",
		})
		return
	}

	user := models.User{Username: body.Username, Password: string(hash)}
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed tocreate user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}
	var user models.User
	db.DB.First(&user, "username = ?", body.Username)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	secretKey, err := utils.ReadEnv("SECRET_KEY")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token(env error)",
		})
		return
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
