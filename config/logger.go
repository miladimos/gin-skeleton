package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Init() {
	customFormatter := new(logger.JSONFormatter)
	logger.SetFormatter(customFormatter)
	logger.SetReportCaller(true)

	// loglevel := "logger.level"
	// SetLogLevel(loglevel)

	// if config.loggin.stdout {
	// 	logger.New().Out = os.Stdout
	// }
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		fmt.Println("failed to log", err.Error())
	}

}

// func SetLogLevel(logLevel string) {
// 	switch strings.ToLower(logLevel) {
// 	case "debug':
// 		logger.SetLevel(logger.DebugLevel)
// 	default:
// 		logger.SetLevel(logger.DebugLevel)
// }

func LogInfo(message string, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("version"),
	}).Info(message)
}

func LogError(message string, err error, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"error":        err.Error(),
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("version"),
	}).Error(message)
}

func LogFatal(message string, err error, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"error":        err.Error(),
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("version"),
	}).Fatal(message)
}

func LogDebug(message string, path error, xRequestID string, errors error) {
	logger.WithFields(logger.Fields{
		"path":         path,
		"error":        "N/A",
		"x-request-id": xRequestID,
		// "version" : c.Request.Header.Get("version")
	}).Debug(message)
}

func PanicLn(message string) {
	logger.Panicln(message)
}
