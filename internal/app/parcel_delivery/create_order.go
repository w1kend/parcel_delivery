package parcel_delivery

import (
	context "context"

	"github.com/google/uuid"
	"github.com/w1kend/guavapay_test/internal/pkg/api/parcel_delivery_grpc"
)

// CreateOrder implements parcel_delivery.ParcelDeliveryServer
func (i *Implementation) CreateOrder(context.Context, *parcel_delivery_grpc.CreateOrderRequest) (*parcel_delivery_grpc.CreateOrderResponse, error) {

	return &parcel_delivery_grpc.CreateOrderResponse{
		Uuid: uuid.New().String(),
	}, nil
}
