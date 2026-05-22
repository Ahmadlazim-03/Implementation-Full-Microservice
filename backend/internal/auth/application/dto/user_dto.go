package dto

import "backend/internal/auth/domain/entity"

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func ToUserResponse(u *entity.User) UserResponse {
	return UserResponse{
		ID:    u.ID(),
		Email: u.Email().String(),
		Name:  u.Name(),
		Role:  string(u.Role()),
	}
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
