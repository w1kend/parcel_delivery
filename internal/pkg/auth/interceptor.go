package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authMetadataName = "authorization"

type AuthInterceptor struct {
	tokenManager TokenManager
	whiteList    map[string]interface{}
}

func NewAuthInterceptor(tokenManager TokenManager, whiteList map[string]interface{}) AuthInterceptor {
	return AuthInterceptor{
		tokenManager: tokenManager,
		whiteList:    whiteList,
	}
}

func (i AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		err = i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i AuthInterceptor) authorize(ctx context.Context, method string) error {
	if _, noAuth := i.whiteList[method]; noAuth {
		return nil
	}

	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return status.Error(codes.Unauthenticated, "there is no metadata")
	}

	values := md.Get(authMetadataName)
	if len(values) == 0 {
		return status.Error(codes.Unauthenticated, "there is no authorization metadata")
	}

	tokenStr := values[0]
	claims, err := i.tokenManager.Parse(tokenStr)
	if err != nil {
		return status.Error(codes.Unauthenticated, "failed to parse token")
	}

	if claims.Valid() != nil {
		return status.Error(codes.Unauthenticated, "token is not valid")
	}

	return nil
}
