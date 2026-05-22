package query

import (
	"context"

	"backend/internal/review/application/dto"
	"backend/internal/review/domain/repository"
)

type ListReviewsByPlaceHandler struct {
	reviews repository.ReviewRepository
}

func NewListReviewsByPlaceHandler(reviews repository.ReviewRepository) *ListReviewsByPlaceHandler {
	return &ListReviewsByPlaceHandler{reviews: reviews}
}

func (h *ListReviewsByPlaceHandler) Handle(ctx context.Context, placeID string) ([]dto.ReviewResponse, error) {
	items, err := h.reviews.FindByPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}
	out := make([]dto.ReviewResponse, 0, len(items))
	for _, r := range items {
		out = append(out, dto.ToReviewResponse(r))
	}
	return out, nil
}
