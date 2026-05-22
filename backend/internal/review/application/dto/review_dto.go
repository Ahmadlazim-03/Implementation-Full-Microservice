package dto

import "backend/internal/review/domain/entity"

type ReviewResponse struct {
	ID      string `json:"id"`
	PlaceID string `json:"place_id"`
	UserID  string `json:"user_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

func ToReviewResponse(r *entity.Review) ReviewResponse {
	return ReviewResponse{
		ID: r.ID(), PlaceID: r.PlaceID(), UserID: r.UserID(),
		Rating: r.Rating(), Comment: r.Comment(),
	}
}
