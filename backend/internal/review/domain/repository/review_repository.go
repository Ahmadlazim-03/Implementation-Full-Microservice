package repository

import (
	"context"

	"backend/internal/review/domain/entity"
)

type ReviewRepository interface {
	Save(ctx context.Context, r *entity.Review) error
	FindByPlace(ctx context.Context, placeID string) ([]*entity.Review, error)
}
