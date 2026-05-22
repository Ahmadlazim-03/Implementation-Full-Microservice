package dto

import "backend/internal/places/domain/entity"

type PlaceResponse struct {
	ID          string  `json:"id"`
	CategoryID  string  `json:"category_id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
}

func ToPlaceResponse(p *entity.Place) PlaceResponse {
	return PlaceResponse{
		ID:          p.ID(),
		CategoryID:  p.CategoryID(),
		Name:        p.Name(),
		Latitude:    p.Location().Latitude(),
		Longitude:   p.Location().Longitude(),
		Address:     p.Address(),
		Description: p.Description(),
	}
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func ToCategoryResponse(c *entity.Category) CategoryResponse {
	return CategoryResponse{ID: c.ID(), Name: c.Name(), Icon: c.Icon()}
}
