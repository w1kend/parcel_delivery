package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenManager struct {
	secret string
}

func NewTokenManager(secret string) TokenManager {
	return TokenManager{
		secret: secret,
	}
}

func (m TokenManager) Sign(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(m.secret))
}

func (m TokenManager) Parse(tokenString string) (*AClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(*AClaims), nil
}
