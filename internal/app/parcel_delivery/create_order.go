package parcel_delivery

import (
	context "context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/api/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateOrder implements parcel_delivery.ParcelDeliveryServer
func (i *Implementation) CreateOrder(ctx context.Context, req *parcel_delivery_grpc.CreateOrderRequest) (*parcel_delivery_grpc.CreateOrderResponse, error) {
	//request validation

	order := model.Orders{
		ID:                uuid.New(),
		FromAddr:          req.GetFrom(),
		ToAddr:            req.GetTo(),
		Status:            model.OrderStatus_New,
		Price:             int16(rand.Int()),
		SenderName:        req.GetSenderName(),
		SenderPassportNum: req.GetSenderPassportNum(),
		RecipientName:     req.GetRecipientName(),
		Weight:            int16(req.GetWeight()),
		CreatedAt:         time.Now(),
		CreatedBy:         uuid.New(),
	}

	err := i.OrdersRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &parcel_delivery_grpc.CreateOrderResponse{
		Uuid: order.ID.String(),
	}, nil
}
