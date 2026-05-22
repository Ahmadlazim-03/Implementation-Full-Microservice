package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
	emailRegex          = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	ErrInvalidEmail     = errors.New("invalid email format")
)

// Email is a Value Object: immutable + self-validating.
type Email struct {
	value string
}

func NewEmail(raw string) (Email, error) {
	v := strings.ToLower(strings.TrimSpace(raw))
	if !emailRegex.MatchString(v) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: v}, nil
}

func (e Email) String() string { return e.value }
func (e Email) Equals(o Email) bool { return e.value == o.value }
