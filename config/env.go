package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
		panic("Error in loading .env file")
	}

	log.Println("Env file loaded ...")

}

func EnvGet(key string) string {
	return os.Getenv(key)
}
