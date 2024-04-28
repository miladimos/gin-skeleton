package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("[%s] %s %s %v", c.Request.Method, c.Request.URL.Path, c.ClientIP(), latency)
	}
}

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

// func LogRequestInfo() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		accessLogMap := make(map[string]string)
// 		accessLogMap["request_time"] = time.Now().String()
// 		accessLogMap["request_method"] = c.Request.Method
// 		accessLogMap["request_uri"] = c.Request.RequestURI
// 		accessLogMap["request_proto"] = c.Request.Proto
// 		accessLogMap["request_ua"] = c.Request.UserAgent()
// 		accessLogMap["request_referer"] = c.Request.Referer()
// 		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
// 		accessLogMap["request_client_ip"] = c.ClientIP()
// 		accessLogMap["x-request_id"] = c.Request.Header.Get("X-Request-Id")

// 		logData, err := json.Marshal(accessLogMap)
// 		if err != nil {
// 			fmt.PrintLn(err.Error())
// 			return
// 		}

// 		accessLogJson := string(logData)
// 		logger.PrintLn(accessLogJson)
// 		c.Next()
// 	}
// }
