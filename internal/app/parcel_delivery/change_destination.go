package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery_test/pkg/parcel_delivery_grpc"
)

// ChangeDestination implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) ChangeDestination(context.Context, *parcel_delivery_grpc.ChangeDestinationRequest) (*parcel_delivery_grpc.ChengeDestinationResponse, error) {
	panic("unimplemented")
}
