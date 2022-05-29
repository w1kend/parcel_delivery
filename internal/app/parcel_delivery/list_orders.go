package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery_test/internal/pkg/api/parcel_delivery_grpc"
)

// ListOrders implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) ListOrders(context.Context, *parcel_delivery_grpc.ListOrdersRequest) (*parcel_delivery_grpc.ListOrdersResponse, error) {
	panic("unimplemented")
}
