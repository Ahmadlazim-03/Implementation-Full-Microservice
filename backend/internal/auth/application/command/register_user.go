package command

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"backend/internal/auth/application/dto"
	"backend/internal/auth/domain/entity"
	"backend/internal/auth/domain/repository"
	"backend/internal/auth/domain/service"
	"backend/internal/auth/domain/valueobject"
	shareddomain "backend/internal/shared/domain"
)

type RegisterUserInput struct {
	Email    string
	Password string
	Name     string
}

// RegisterUserHandler is an application service (use case).
// It orchestrates domain objects — no business rules live here.
type RegisterUserHandler struct {
	users  repository.UserRepository
	hasher service.PasswordHasher
	tokens service.TokenService
}

func NewRegisterUserHandler(users repository.UserRepository, hasher service.PasswordHasher, tokens service.TokenService) *RegisterUserHandler {
	return &RegisterUserHandler{users: users, hasher: hasher, tokens: tokens}
}

func (h *RegisterUserHandler) Handle(ctx context.Context, in RegisterUserInput) (*dto.AuthResponse, error) {
	email, err := valueobject.NewEmail(in.Email)
	if err != nil {
		return nil, err
	}
	if len(in.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	exists, err := h.users.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, shareddomain.ErrAlreadyExists
	}

	hash, err := h.hasher.Hash(in.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(uuid.NewString(), email, hash, in.Name, entity.RoleUser)
	if err := h.users.Save(ctx, user); err != nil {
		return nil, err
	}

	token, err := h.tokens.Issue(user)
	if err != nil {
		return nil, err
	}

	resp := &dto.AuthResponse{Token: token, User: dto.ToUserResponse(user)}
	return resp, nil
}
