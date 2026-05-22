package command

import (
	"context"

	"github.com/google/uuid"

	"backend/internal/places/application/dto"
	"backend/internal/places/domain/entity"
	"backend/internal/places/domain/repository"
	"backend/internal/shared/valueobject"
)

type CreatePlaceInput struct {
	CategoryID  string
	Name        string
	Latitude    float64
	Longitude   float64
	Address     string
	Description string
}

type CreatePlaceHandler struct {
	places repository.PlaceRepository
}

func NewCreatePlaceHandler(places repository.PlaceRepository) *CreatePlaceHandler {
	return &CreatePlaceHandler{places: places}
}

func (h *CreatePlaceHandler) Handle(ctx context.Context, in CreatePlaceInput) (*dto.PlaceResponse, error) {
	coord, err := valueobject.NewCoordinate(in.Latitude, in.Longitude)
	if err != nil {
		return nil, err
	}
	place, err := entity.NewPlace(uuid.NewString(), in.CategoryID, in.Name, coord, in.Address, in.Description)
	if err != nil {
		return nil, err
	}
	if err := h.places.Save(ctx, place); err != nil {
		return nil, err
	}
	resp := dto.ToPlaceResponse(place)
	return &resp, nil
}
