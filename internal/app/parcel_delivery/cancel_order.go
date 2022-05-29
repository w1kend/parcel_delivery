package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery_test/internal/pkg/api/parcel_delivery_grpc"
)

// CancelOrder implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) CancelOrder(context.Context, *parcel_delivery_grpc.CancelOrderRequest) (*parcel_delivery_grpc.CancelOrderResponse, error) {
	panic("unimplemented")
}
