package query

import (
	"context"

	"backend/internal/places/application/dto"
	"backend/internal/places/domain/repository"
	"backend/internal/shared/valueobject"
)

type FindNearbyPlacesInput struct {
	Latitude     float64
	Longitude    float64
	RadiusMeters float64
	Limit        int64
}

type FindNearbyPlacesHandler struct {
	places repository.PlaceRepository
}

func NewFindNearbyPlacesHandler(places repository.PlaceRepository) *FindNearbyPlacesHandler {
	return &FindNearbyPlacesHandler{places: places}
}

func (h *FindNearbyPlacesHandler) Handle(ctx context.Context, in FindNearbyPlacesInput) ([]dto.PlaceResponse, error) {
	center, err := valueobject.NewCoordinate(in.Latitude, in.Longitude)
	if err != nil {
		return nil, err
	}
	if in.RadiusMeters <= 0 {
		in.RadiusMeters = 1000 // default 1 km
	}
	items, err := h.places.FindNearby(ctx, center, in.RadiusMeters, in.Limit)
	if err != nil {
		return nil, err
	}
	out := make([]dto.PlaceResponse, 0, len(items))
	for _, p := range items {
		out = append(out, dto.ToPlaceResponse(p))
	}
	return out, nil
}
