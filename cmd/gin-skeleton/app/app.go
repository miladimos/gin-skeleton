package app

import "gin-skeleton/internal/routes"

func Run() {
	router := routes.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
