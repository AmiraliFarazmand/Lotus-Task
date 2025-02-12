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

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check if the token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// find expected User
		var user models.User
		db.DB.First(&user, claims["sub"])
		// check if the user exists
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// set the user to the context
		c.Set("user", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
