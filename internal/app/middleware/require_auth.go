package middleware

import (
	"fmt"
	"log"
	"lotus-task/internal/app/db"
	"lotus-task/internal/app/models"
	"lotus-task/internal/app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func validateClaims(claims jwt.MapClaims) (models.User, bool) {
	// check if the token is expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return models.User{}, false
	}

	// find expected User
	var user models.User
	db.DB.First(&user, claims["sub"])

	// check if the user exists
	if user.ID == 0 {
		return models.User{}, false
	}
	return user, true
}

func RequireAuth(c *gin.Context) {
	// get the cookie off the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey, _ := utils.ReadEnv("SECRET_KEY")
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, claimsOk := token.Claims.(jwt.MapClaims)
	if !claimsOk {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, ok := validateClaims(claims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// set the user to the context
	c.Set("user", user)
	c.Next()
}
