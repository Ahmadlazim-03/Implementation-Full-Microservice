package service

import "backend/internal/auth/domain/entity"

// TokenService issues and validates auth tokens.
// Implemented in infrastructure (e.g., JWT).
type TokenService interface {
	Issue(user *entity.User) (string, error)
	Verify(token string) (Claims, error)
}

type Claims struct {
	UserID string
	Email  string
	Role   string
}
