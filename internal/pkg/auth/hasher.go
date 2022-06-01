package auth

import (
	"golang.org/x/crypto/bcrypt"
)

//
type Hasher struct {
	cost int
}

//
func NewHasher(cost int) Hasher {
	return Hasher{
		cost: cost,
	}
}

//
func (h Hasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (h Hasher) IsValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
