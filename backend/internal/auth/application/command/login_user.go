package command

import (
	"context"
	"errors"

	"backend/internal/auth/application/dto"
	"backend/internal/auth/domain/repository"
	"backend/internal/auth/domain/service"
	"backend/internal/auth/domain/valueobject"
	shareddomain "backend/internal/shared/domain"
)

type LoginUserInput struct {
	Email    string
	Password string
}

type LoginUserHandler struct {
	users  repository.UserRepository
	hasher service.PasswordHasher
	tokens service.TokenService
}

func NewLoginUserHandler(users repository.UserRepository, hasher service.PasswordHasher, tokens service.TokenService) *LoginUserHandler {
	return &LoginUserHandler{users: users, hasher: hasher, tokens: tokens}
}

func (h *LoginUserHandler) Handle(ctx context.Context, in LoginUserInput) (*dto.AuthResponse, error) {
	email, err := valueobject.NewEmail(in.Email)
	if err != nil {
		return nil, err
	}

	user, err := h.users.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, shareddomain.ErrNotFound) {
			return nil, shareddomain.ErrUnauthorized
		}
		return nil, err
	}

	if err := h.hasher.Verify(in.Password, user.PasswordHash()); err != nil {
		return nil, shareddomain.ErrUnauthorized
	}

	token, err := h.tokens.Issue(user)
	if err != nil {
		return nil, err
	}
	return &dto.AuthResponse{Token: token, User: dto.ToUserResponse(user)}, nil
}
