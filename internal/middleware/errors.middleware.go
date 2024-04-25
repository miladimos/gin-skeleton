package middleware

import (
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
