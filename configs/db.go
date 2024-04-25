package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	dsn := fmt.Sprintf("postgres://pg:pass@localhost:5432/crud",
		os.Getenv("DB_USER"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
		log.Fatalln(err)
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		log.Fatalln(err)
	}

	return database
}
