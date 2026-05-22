package valueobject

import "errors"

// Coordinate is a Value Object representing a geographic point.
// Immutable — create a new one to "change" it.
type Coordinate struct {
	latitude  float64
	longitude float64
}

var ErrInvalidCoordinate = errors.New("invalid coordinate: latitude must be [-90,90], longitude must be [-180,180]")

func NewCoordinate(lat, lng float64) (Coordinate, error) {
	if lat < -90 || lat > 90 || lng < -180 || lng > 180 {
		return Coordinate{}, ErrInvalidCoordinate
	}
	return Coordinate{latitude: lat, longitude: lng}, nil
}

func (c Coordinate) Latitude() float64  { return c.latitude }
func (c Coordinate) Longitude() float64 { return c.longitude }

func (c Coordinate) Equals(other Coordinate) bool {
	return c.latitude == other.latitude && c.longitude == other.longitude
}
