package parcel_delivery

import (
	context "context"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/api/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetOrder implements parcel_delivery.ParcelDeliveryServer
func (i *Implementation) GetOrder(ctx context.Context, req *parcel_delivery_grpc.GetOrderRequest) (*parcel_delivery_grpc.GetOrderResponse, error) {
	order, err := i.OrdersRepo.GetOrder(ctx, req.GetUuid())
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
		Status:        "",
		CreatedAt:     "",
		Price:         1,
	}, nil
}
