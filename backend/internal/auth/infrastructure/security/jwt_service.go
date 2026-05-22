package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"backend/internal/auth/domain/entity"
	"backend/internal/auth/domain/service"
)

type JWTService struct {
	secret      []byte
	expiresHours int
}

func NewJWTService(secret string, expiresHours int) *JWTService {
	return &JWTService{secret: []byte(secret), expiresHours: expiresHours}
}

type jwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func (s *JWTService) Issue(user *entity.User) (string, error) {
	claims := jwtClaims{
		Email: user.Email().String(),
		Role:  string(user.Role()),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.expiresHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.secret)
}

func (s *JWTService) Verify(token string) (service.Claims, error) {
	var c jwtClaims
	parsed, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})
	if err != nil || !parsed.Valid {
		return service.Claims{}, errors.New("invalid token")
	}
	return service.Claims{UserID: c.Subject, Email: c.Email, Role: c.Role}, nil
}
