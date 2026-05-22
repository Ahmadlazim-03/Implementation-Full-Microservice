package entity

import (
	"errors"
	"time"

	"backend/internal/shared/valueobject"
)

// Place is the aggregate root of the Places bounded context.
type Place struct {
	id          string
	categoryID  string
	name        string
	location    valueobject.Coordinate
	address     string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

var ErrInvalidPlace = errors.New("invalid place: name and category are required")

func NewPlace(id, categoryID, name string, location valueobject.Coordinate, address, description string) (*Place, error) {
	if name == "" || categoryID == "" {
		return nil, ErrInvalidPlace
	}
	now := time.Now().UTC()
	return &Place{
		id:          id,
		categoryID:  categoryID,
		name:        name,
		location:    location,
		address:     address,
		description: description,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

func Hydrate(id, categoryID, name string, location valueobject.Coordinate, address, description string, createdAt, updatedAt time.Time) *Place {
	return &Place{
		id: id, categoryID: categoryID, name: name, location: location,
		address: address, description: description, createdAt: createdAt, updatedAt: updatedAt,
	}
}

func (p *Place) ID() string                       { return p.id }
func (p *Place) CategoryID() string               { return p.categoryID }
func (p *Place) Name() string                     { return p.name }
func (p *Place) Location() valueobject.Coordinate { return p.location }
func (p *Place) Address() string                  { return p.address }
func (p *Place) Description() string              { return p.description }
func (p *Place) CreatedAt() time.Time             { return p.createdAt }
func (p *Place) UpdatedAt() time.Time             { return p.updatedAt }
