package parcel_delivery

import (
	context "context"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/w1kend/parcel_delivery/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery/internal/pkg/conv"
	"github.com/w1kend/parcel_delivery/internal/pkg/repositories"
	"github.com/w1kend/parcel_delivery/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetOrder(ctx context.Context, req *parcel_delivery_grpc.GetOrderRequest) (*parcel_delivery_grpc.GetOrderResponse, error) {
	//request validation

	user := auth.UserInfoFromContext(ctx)
	if user == nil {
		return nil, status.Error(codes.Internal, "failed to get user credentials")
	}

	order, err := i.OrdersRepo.GetOrder(ctx, repositories.OrdersFilter{
		OrderID:   conv.Pointer(req.GetUuid()),
		CreatedBy: &user.UserID,
	})
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return &parcel_delivery_grpc.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &parcel_delivery_grpc.GetOrderResponse{
		From:          order.FromAddr,
		To:            order.ToAddr,
		SenderName:    order.SenderName,
		RecipientName: order.RecipientName,
		Weight:        int64(order.Weight),
		Status:        order.Status.String(),
		CreatedAt:     order.CreatedAt.String(),
		Price:         uint64(order.Price),
	}, nil
}
