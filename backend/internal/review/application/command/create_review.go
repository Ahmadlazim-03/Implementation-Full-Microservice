package command

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"backend/internal/review/application/dto"
	"backend/internal/review/domain/entity"
	"backend/internal/review/domain/repository"
	"backend/internal/review/domain/service"
	shareddomain "backend/internal/shared/domain"
)

type CreateReviewInput struct {
	PlaceID string
	UserID  string
	Rating  int
	Comment string
}

type CreateReviewHandler struct {
	reviews  repository.ReviewRepository
	verifier service.PlaceVerifier // injected via gRPC adapter
}

func NewCreateReviewHandler(reviews repository.ReviewRepository, verifier service.PlaceVerifier) *CreateReviewHandler {
	return &CreateReviewHandler{reviews: reviews, verifier: verifier}
}

var ErrPlaceNotFound = errors.New("place not found")

func (h *CreateReviewHandler) Handle(ctx context.Context, in CreateReviewInput) (*dto.ReviewResponse, error) {
	// Cross-context invariant: review hanya untuk place yang valid.
	// Verifikasi via internal gRPC (protobuf) ke places-service.
	exists, err := h.verifier.Exists(ctx, in.PlaceID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, shareddomain.ErrNotFound
	}

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
