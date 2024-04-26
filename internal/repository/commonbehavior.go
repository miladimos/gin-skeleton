package repository

import (
	"context"
	"gin-skeleton/internal/models"

	"gorm.io/gorm"
)

type commonBehavior[T models.DBModel] struct {
	db *gorm.DB
}

// ByField implements CommonBehaviorRepository.
func (*commonBehavior[T]) ByField(ctx context.Context, field string, id uint) (T, error) {
	panic("unimplemented")
}

// ByID implements CommonBehaviorRepository.
func (*commonBehavior[T]) ByID(ctx context.Context, id uint) (T, error) {
	panic("unimplemented")
}

// Delete implements CommonBehaviorRepository.
func (*commonBehavior[T]) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// Save implements CommonBehaviorRepository.
func (*commonBehavior[T]) Save(ctx context.Context, model T) error {
	panic("unimplemented")
}

// Update implements CommonBehaviorRepository.
func (*commonBehavior[T]) Update(ctx context.Context, model T) error {
	panic("unimplemented")
}

func NewCommonBehavior[T models.DBModel](db *gorm.DB) CommonBehaviorRepository[T] {
	return &commonBehavior[T]{
		db: db,
	}
}
