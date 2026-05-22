package security

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{cost: bcrypt.DefaultCost}
}

func (b *BcryptHasher) Hash(plain string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plain), b.cost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

func (b *BcryptHasher) Verify(plain, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
