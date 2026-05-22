package query

import (
	"context"

	"backend/internal/places/application/dto"
	"backend/internal/places/domain/repository"
)

type ListPlacesInput struct {
	CategoryID string
	Search     string
	Limit      int64
	Skip       int64
}

type ListPlacesHandler struct {
	places repository.PlaceRepository
}

func NewListPlacesHandler(places repository.PlaceRepository) *ListPlacesHandler {
	return &ListPlacesHandler{places: places}
}

func (h *ListPlacesHandler) Handle(ctx context.Context, in ListPlacesInput) ([]dto.PlaceResponse, error) {
	if in.Limit <= 0 {
		in.Limit = 50
	}
	items, err := h.places.FindAll(ctx, repository.PlaceFilter{
		CategoryID: in.CategoryID,
		Search:     in.Search,
		Limit:      in.Limit,
		Skip:       in.Skip,
	})
	if err != nil {
		return nil, err
	}
	out := make([]dto.PlaceResponse, 0, len(items))
	for _, p := range items {
		out = append(out, dto.ToPlaceResponse(p))
	}
	return out, nil
}
