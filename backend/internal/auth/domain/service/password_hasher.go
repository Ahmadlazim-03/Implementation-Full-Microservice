package service

// PasswordHasher is a domain port for password operations.
// Implemented in infrastructure (e.g., bcrypt).
type PasswordHasher interface {
	Hash(plain string) (string, error)
	Verify(plain, hash string) error
}
