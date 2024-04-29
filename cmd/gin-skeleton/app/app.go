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
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
