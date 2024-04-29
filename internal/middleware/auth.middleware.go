package middleware

import (
	"fmt"
	"gin-skeleton/config"
	"gin-skeleton/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			// 	c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})

			return
		}

		// Verify the token
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()

		// Check if the user is authenticated
		if isAuthenticated(c) {
			c.Next()
			return
		}
		// User is not authenticated, return an error response
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func isAuthenticated(c *gin.Context) bool {
	// Check if the user is authenticated based on a JWT token, session, or any other mechanism
	// Return true if the user is authenticated, false otherwise
	return true
}

func AuthCookie(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		config.DB.Find(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
