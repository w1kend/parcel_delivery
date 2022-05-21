package parcel_delivery

import (
	context "context"

	"github.com/w1kend/guavapay_test/internal/pkg/api/parcel_delivery_grpc"
)

// ChangeDestination implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) ChangeDestination(context.Context, *parcel_delivery_grpc.ChangeDestinationRequest) (*parcel_delivery_grpc.ChengeDestinationResponse, error) {
	panic("unimplemented")
}
