package main

import (
	"gin-skeleton/config"
	"gin-skeleton/internal/models"
)

func init() {
	config.SetupEnvVariables()
	config.SetupDatabase()
}

func main() {
	config.DB.AutoMigrate(&models.Tag{})
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Post{})

	// if err := database.AutoMigrate(&models.User{}); err != nil {
	// 	panic(err)
	// }
}
