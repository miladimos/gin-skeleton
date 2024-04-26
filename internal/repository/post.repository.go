package repository

import (
	"gorm.io/gorm"
)

type PostRepository interface {
}

type postRepository struct {
	connection *gorm.DB
}

func NewPostRepository(connection *gorm.DB) PostRepository {
	return &postRepository{
		connection: connection,
	}
}
