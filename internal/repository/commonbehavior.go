package repository

import (
	"context"
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

func (c *commonBehavior[T]) ByID(ctx context.Context, id uint) (T, error) {
	return c.ByField(ctx, "id", id)
}

func (c *commonBehavior[T]) ByField(ctx context.Context, field string, id uint) (T, error) {
	var t T
	err := c.db.WithContext(ctx).Model(t).Where(field+"=?", id).First(t).Error
	return t, err
}

func (c *commonBehavior[T]) Save(ctx context.Context, model T) error {
	return c.db.WithContext(ctx).Save(model).Error
}
