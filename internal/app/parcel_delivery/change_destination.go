package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery/internal/pkg/conv"
	"github.com/w1kend/parcel_delivery/internal/pkg/repositories"
	"github.com/w1kend/parcel_delivery/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ChangeDestination(ctx context.Context, req *parcel_delivery_grpc.ChangeDestinationRequest) (*parcel_delivery_grpc.ChengeDestinationResponse, error) {
	//request validation

	user := auth.UserInfoFromContext(ctx)
	if user == nil {
		return nil, status.Error(codes.Internal, "failed to get user credentials")
	}

	order, err := i.OrdersRepo.GetOrder(ctx, repositories.OrdersFilter{
		OrderID:   conv.Pointer(req.GetUuid()),
		CreatedBy: &user.UserID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if order == nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	if order.Status != model.OrderStatus_New {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = i.OrdersRepo.ChangeDestination(ctx, req.GetUuid(), req.GetNewDestination())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &parcel_delivery_grpc.ChengeDestinationResponse{}, nil
}
