package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/model"
)

const (
	TokenTTLHours = time.Hour * 24
)

type AClaims struct {
	jwt.StandardClaims
	UserID string         `json:"user_id"`
	Role   model.UserRole `json:"role"`
}

type UserInfo struct {
	UserID string
	Role   model.UserRole
}

func UserInfoFromContext(ctx context.Context) *UserInfo {
	user := ctx.Value(CtxUserInfo).(UserInfo)
	if user.UserID == "" {
		return nil
	}

	return &user
}

func NewClaims(userID string, role model.UserRole) AClaims {
	return AClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTLHours).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}
