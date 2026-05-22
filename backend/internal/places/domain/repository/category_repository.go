package repository

import (
	"context"

	"backend/internal/places/domain/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, c *entity.Category) error
	FindByID(ctx context.Context, id string) (*entity.Category, error)
	FindAll(ctx context.Context) ([]*entity.Category, error)
}
