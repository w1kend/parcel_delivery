package parcel_delivery

import (
	context "context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery_test/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SignUp implements parcel_delivery.ParcelDeliveryServer
func (i *Implementation) SignUp(ctx context.Context, req *parcel_delivery_grpc.SignUpRequest) (*parcel_delivery_grpc.SignUpResponse, error) {
	//validation

	hashedPassword, err := i.Hasher.Hash(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := model.Users{
		ID:        uuid.New(),
		Name:      "user_name",
		Email:     req.GetEmail(),
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	err = i.UsersRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := i.TokenManager.Sign(auth.NewClaims(user.ID.String(), "role"))
	if err != nil {
		fmt.Println(err)
	}

	return &parcel_delivery_grpc.SignUpResponse{
		Token: token,
	}, nil
}
