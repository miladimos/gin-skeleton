package app

import (
	logger "gin-skeleton/configs/logger"
	"gin-skeleton/internal/routes"
)

func Run() {

	// configs.SetupDatabase()
	logger.Init()

	router := routes.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
