package entity

import (
	"time"

	"backend/internal/auth/domain/valueobject"
)

// Role is a domain enum for the user's permission level.
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

// User is the aggregate root of the Auth bounded context.
// All invariants live here — repository must never bypass these methods.
type User struct {
	id           string
	email        valueobject.Email
	passwordHash string
	name         string
	role         Role
	createdAt    time.Time
	updatedAt    time.Time
}

func NewUser(id string, email valueobject.Email, passwordHash, name string, role Role) *User {
	now := time.Now().UTC()
	return &User{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		name:         name,
		role:         role,
		createdAt:    now,
		updatedAt:    now,
	}
}

// Hydrate rebuilds a User from persistence without re-running invariants.
// Only repositories should call this.
func Hydrate(id string, email valueobject.Email, passwordHash, name string, role Role, createdAt, updatedAt time.Time) *User {
	return &User{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		name:         name,
		role:         role,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

func (u *User) ID() string                   { return u.id }
func (u *User) Email() valueobject.Email     { return u.email }
func (u *User) PasswordHash() string         { return u.passwordHash }
func (u *User) Name() string                 { return u.name }
func (u *User) Role() Role                   { return u.role }
func (u *User) CreatedAt() time.Time         { return u.createdAt }
func (u *User) UpdatedAt() time.Time         { return u.updatedAt }
func (u *User) IsAdmin() bool                { return u.role == RoleAdmin }

func (u *User) Rename(name string) {
	u.name = name
	u.updatedAt = time.Now().UTC()
}
