package repository

import (
	"context"

	"backend/internal/auth/domain/entity"
	"backend/internal/auth/domain/valueobject"
)

// UserRepository is a domain port. Infrastructure implements it.
// The domain layer must never import infrastructure.
type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error)
	ExistsByEmail(ctx context.Context, email valueobject.Email) (bool, error)
}
