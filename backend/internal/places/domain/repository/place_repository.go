package repository

import (
	"context"

	"backend/internal/places/domain/entity"
)

type PlaceRepository interface {
	Save(ctx context.Context, p *entity.Place) error
	FindByID(ctx context.Context, id string) (*entity.Place, error)
	FindAll(ctx context.Context, filter PlaceFilter) ([]*entity.Place, error)
}

type PlaceFilter struct {
	CategoryID string
	Search     string
	Limit      int64
	Skip       int64
}
