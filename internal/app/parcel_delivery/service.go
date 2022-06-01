package parcel_delivery

import (
	"github.com/w1kend/parcel_delivery_test/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery_test/internal/repositories"
	"github.com/w1kend/parcel_delivery_test/pkg/parcel_delivery_grpc"
)

type Implementation struct {
	parcel_delivery_grpc.UnimplementedParcelDeliveryServer
	OrdersRepo   repositories.Orders
	UsersRepo    repositories.Users
	Hasher       auth.Hasher
	TokenManager auth.TokenManager
}

func NewImplementation(
	ordersRepo repositories.Orders,
	usersRepo repositories.Users,
	hasher auth.Hasher,
	tokenManager auth.TokenManager,
) Implementation {
	return Implementation{
		OrdersRepo:   ordersRepo,
		UsersRepo:    usersRepo,
		Hasher:       hasher,
		TokenManager: tokenManager,
	}
}

// var _ parcel_delivery.ParcelDeliveryServer = &Implementation{}
