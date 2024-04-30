package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "True")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Catch-Control, X-Request-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
// }



// c.Writer.Header().Set("Content-Type", "application/json")
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(200)
// 	} else {
// 		c.Next()
// 	}
