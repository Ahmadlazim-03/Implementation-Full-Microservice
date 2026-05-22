package entity

import (
	"errors"
	"time"
)

var ErrInvalidRating = errors.New("rating must be between 1 and 5")

type Review struct {
	id        string
	placeID   string
	userID    string
	rating    int
	comment   string
	createdAt time.Time
}

func NewReview(id, placeID, userID string, rating int, comment string) (*Review, error) {
	if rating < 1 || rating > 5 {
		return nil, ErrInvalidRating
	}
	if placeID == "" || userID == "" {
		return nil, errors.New("placeID and userID are required")
	}
	return &Review{
		id: id, placeID: placeID, userID: userID,
		rating: rating, comment: comment,
		createdAt: time.Now().UTC(),
	}, nil
}

func Hydrate(id, placeID, userID string, rating int, comment string, createdAt time.Time) *Review {
	return &Review{id: id, placeID: placeID, userID: userID, rating: rating, comment: comment, createdAt: createdAt}
}

func (r *Review) ID() string           { return r.id }
func (r *Review) PlaceID() string      { return r.placeID }
func (r *Review) UserID() string       { return r.userID }
func (r *Review) Rating() int          { return r.rating }
func (r *Review) Comment() string      { return r.comment }
func (r *Review) CreatedAt() time.Time { return r.createdAt }
