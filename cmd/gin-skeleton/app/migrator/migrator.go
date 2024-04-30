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

	if err := config.DB.AutoMigrate(
		&models.Tag{},
		&models.User{},
		&models.Post{},
	); err != nil {
		panic(err)
	}
}
