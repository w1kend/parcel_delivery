package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TokenTTLHours = time.Hour * 24
)

type AClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func NewClaims(userID, role string) AClaims {
	return AClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTLHours).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}
