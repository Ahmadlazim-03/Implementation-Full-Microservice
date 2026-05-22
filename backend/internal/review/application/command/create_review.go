package command

import (
	"context"

	"github.com/google/uuid"

	"backend/internal/review/application/dto"
	"backend/internal/review/domain/entity"
	"backend/internal/review/domain/repository"
)

type CreateReviewInput struct {
	PlaceID string
	UserID  string
	Rating  int
	Comment string
}

type CreateReviewHandler struct {
	reviews repository.ReviewRepository
}

func NewCreateReviewHandler(reviews repository.ReviewRepository) *CreateReviewHandler {
	return &CreateReviewHandler{reviews: reviews}
}

func (h *CreateReviewHandler) Handle(ctx context.Context, in CreateReviewInput) (*dto.ReviewResponse, error) {
	r, err := entity.NewReview(uuid.NewString(), in.PlaceID, in.UserID, in.Rating, in.Comment)
	if err != nil {
		return nil, err
	}
	if err := h.reviews.Save(ctx, r); err != nil {
		return nil, err
	}
	resp := dto.ToReviewResponse(r)
	return &resp, nil
}
