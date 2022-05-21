package parcel_delivery

import (
	context "context"

	"github.com/w1kend/guavapay_test/internal/pkg/api/parcel_delivery_grpc"
)

// SignUp implements parcel_delivery.ParcelDeliveryServer
func (*Implementation) SignUp(context.Context, *parcel_delivery_grpc.SignUpRequest) (*parcel_delivery_grpc.SignUpResponse, error) {
	panic("unimplemented")
}
