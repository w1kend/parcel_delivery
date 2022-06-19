package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/repositories"
	"github.com/w1kend/parcel_delivery_test/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListOrders(ctx context.Context, req *parcel_delivery_grpc.ListOrdersRequest) (*parcel_delivery_grpc.ListOrdersResponse, error) {
	//request validation

	user := auth.UserInfoFromContext(ctx)
	if user == nil {
		return nil, status.Error(codes.Internal, "failed to get user credentials")
	}

	orders, err := i.OrdersRepo.ListOrders(ctx, repositories.OrdersFilter{
		CreatedBy: &user.UserID,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return i.toProtoListOrders(orders), nil
}

func (i *Implementation) toProtoListOrders(orders []model.Orders) *parcel_delivery_grpc.ListOrdersResponse {
	protoOrders := make([]*parcel_delivery_grpc.OrderShortInfo, 0, len(orders))

	for _, order := range orders {
		protoOrders = append(protoOrders, &parcel_delivery_grpc.OrderShortInfo{
			Uuid:          order.ID.String(),
			From:          order.FromAddr,
			To:            order.ToAddr,
			RecipientName: order.RecipientName,
			Status:        order.Status.String(),
			CreatedAt:     order.CreatedAt.String(),
		})
	}

	return &parcel_delivery_grpc.ListOrdersResponse{
		Orders: protoOrders,
	}
}
