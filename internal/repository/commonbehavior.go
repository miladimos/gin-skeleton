package repository

import (
	"gin-skeleton/internal/models"

	"gorm.io/gorm"
)

type commonBehavior[T models.DBModel] struct {
	db *gorm.DB
}

func NewCommonBehavior[T models.DBModel](db *gorm.DB) CommonBehaviorRepository[T] {
	return &commonBehavior[T]{
		db: db,
	}
}
