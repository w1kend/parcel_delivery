package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery_test/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery_test/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SignIn(ctx context.Context, req *parcel_delivery_grpc.SignInRequest) (*parcel_delivery_grpc.SignInResponse, error) {
	//validation

	user, err := i.UsersRepo.GetByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil || !i.Hasher.IsValid(req.GetPassword(), user.Password) {
		return nil, status.Error(codes.Unauthenticated, "invalid email or password")
	}

	token, err := i.TokenManager.Sign(auth.NewClaims(user.ID.String(), "role"))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &parcel_delivery_grpc.SignInResponse{
		Token: token,
	}, nil
}
