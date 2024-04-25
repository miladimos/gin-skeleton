package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// before request
		context.Next()
		// after request

		// start := time.Now()
		// c.Next()
		// duration := time.Since(start)
		// log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}
