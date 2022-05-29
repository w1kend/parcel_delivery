package parcel_delivery

import (
	context "context"

	"github.com/w1kend/parcel_delivery_test/internal/pkg/api/parcel_delivery_grpc"
)

// SignIn implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) SignIn(context.Context, *parcel_delivery_grpc.SignInRequest) (*parcel_delivery_grpc.SignInResponse, error) {
	panic("unimplemented")
}
