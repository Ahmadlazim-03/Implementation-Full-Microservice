package repository

import (
	"context"

	"backend/internal/places/domain/entity"
	"backend/internal/shared/valueobject"
)

type PlaceRepository interface {
	Save(ctx context.Context, p *entity.Place) error
	FindByID(ctx context.Context, id string) (*entity.Place, error)
	FindAll(ctx context.Context, filter PlaceFilter) ([]*entity.Place, error)

	// FindNearby returns places within radiusMeters of the given point,
	// ordered by distance ascending. Backed by MongoDB 2dsphere index.
	FindNearby(ctx context.Context, center valueobject.Coordinate, radiusMeters float64, limit int64) ([]*entity.Place, error)
}

type PlaceFilter struct {
	CategoryID string
	Search     string
	Limit      int64
	Skip       int64
}
