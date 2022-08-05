package parcel_delivery

import (
	context "context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/w1kend/parcel_delivery/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateOrder(ctx context.Context, req *parcel_delivery_grpc.CreateOrderRequest) (*parcel_delivery_grpc.CreateOrderResponse, error) {
	//request validation

	user := auth.UserInfoFromContext(ctx)
	if user == nil {
		return nil, status.Error(codes.Internal, "failed to get user credentials")
	}

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
		CreatedBy:         uuid.MustParse(user.UserID),
	}

	err := i.OrdersRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &parcel_delivery_grpc.CreateOrderResponse{
		Uuid: order.ID.String(),
	}, nil
}
