package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// apiKey := c.GetHeader("X-API-Key")
		// if apiKey == "" {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		// 	return
		// }
		// c.Next()

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
