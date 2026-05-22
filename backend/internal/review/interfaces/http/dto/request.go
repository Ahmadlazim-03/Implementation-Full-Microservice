package dto

type CreateReviewRequest struct {
	PlaceID string `json:"place_id" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}
