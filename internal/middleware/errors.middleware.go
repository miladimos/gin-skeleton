package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(404, gin.H{"message": "method not allowed"})
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(404, gin.H{"message": "route not defined"})
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
