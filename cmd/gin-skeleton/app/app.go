package app

import (
	config "gin-skeleton/config"
	"gin-skeleton/internal/routes"
)

// logger "gin-skeleton/configs/logger"

func Run() {

	config.SetupEnvVariables()
	config.SetupDatabase()

	// logger.Init()

	router := routes.SetupRouter()

	port := config.EnvGet("APP_PORT")
	if port == "" {
		port = ":8080"
	}
	err := router.Run(":%d", port)
	if err != nil {
		panic(err)
	}
}
