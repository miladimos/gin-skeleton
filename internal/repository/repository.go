package repository

import (
	"context"
	"gin-skeleton/internal/models"
)

type CommonBehaviorRepository[T models.DBModel] interface {
	ByID(ctx context.Context, id uint) (T, error)
	ByField(ctx context.Context, field string, id uint) (T, error)
	Save(ctx context.Context, model T) error
}
