package controllers

import (
	"net/http"
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/models"
	"lotus-task/internal/app/utils"
	"lotus-task/internal/app/validators"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const tokenExpireTime int = 72 // expire time in hour

type authRequest struct {
	Username string
	Password string
}

func createToken(userID uint) (string, error) {

	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenExpireTime)).Unix(),
	})
	secretKey, err := utils.ReadEnv("SECRET_KEY")
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Signup(c *gin.Context) {
	var body authRequest
	// Validate format of request
	if c.ShouldBindJSON(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Validate Username and Password of request body
	if err := validators.ValidateUsernamePassword(body.Username, body.Password, db.DB); err != nil { 
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Insert into DB
	user := models.User{Username: body.Username, Password: string(hash)}
	result := db.DB.Create(&user)
	if result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func Login(c *gin.Context) {
	var body authRequest
	// Validate format of request
	if c.ShouldBindJSON(&body) != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	var user models.User
	db.DB.First(&user, "username = ?", body.Username)
	if user.ID == 0 {
		utils.RespondWithError(c, http.StatusBadRequest, "User not found")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid username or password")
		return
	}

	tokenString, err := createToken(uint(user.ID))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create token")
		return
	}

	// Set the token in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*tokenExpireTime, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func ValidateIsAuthenticated(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "UnAuthorized user")
		return
	}

	username := user.(models.User).Username
	c.JSON(http.StatusOK, gin.H{
		"message":  "I am Authenticated",
		"username": username,
	})
}
