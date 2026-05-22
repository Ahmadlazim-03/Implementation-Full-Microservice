package query

import (
	"context"

	"backend/internal/places/application/dto"
	"backend/internal/places/domain/repository"
)

type GetPlaceHandler struct {
	places repository.PlaceRepository
}

func NewGetPlaceHandler(places repository.PlaceRepository) *GetPlaceHandler {
	return &GetPlaceHandler{places: places}
}

func (h *GetPlaceHandler) Handle(ctx context.Context, id string) (*dto.PlaceResponse, error) {
	p, err := h.places.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := dto.ToPlaceResponse(p)
	return &resp, nil
}
