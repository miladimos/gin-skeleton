package repository

import (
	"gorm.io/gorm"
)

type AuthRepository interface {
}

type authRepository struct {
	connection *gorm.DB
}

func NewAuthRepository(connection *gorm.DB) AuthRepository {
	return &authRepository{
		connection: connection,
	}
}
