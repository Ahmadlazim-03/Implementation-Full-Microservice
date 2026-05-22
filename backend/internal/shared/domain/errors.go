package shareddomain

import "errors"

// Common domain errors shared across bounded contexts.
var (
	ErrNotFound       = errors.New("entity not found")
	ErrAlreadyExists  = errors.New("entity already exists")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
)
