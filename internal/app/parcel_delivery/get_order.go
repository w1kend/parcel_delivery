package parcel_delivery

import (
	context "context"

	"github.com/w1kend/guavapay_test/internal/pkg/api/parcel_delivery_grpc"
)

// GetOrder implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) GetOrder(context.Context, *parcel_delivery_grpc.GetOrderRequest) (*parcel_delivery_grpc.GetOrderResponse, error) {
	panic("unimplemented")
}
