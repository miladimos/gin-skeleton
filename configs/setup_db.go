package configs

import (
	"fmt"
	"os"

	"gin-skeleton/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("postgres://pg:pass@localhost:5432/crud",
		os.Getenv("DB_USER"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	if err := database.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	return database
}
