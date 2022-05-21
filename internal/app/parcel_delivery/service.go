package parcel_delivery

import (
	"database/sql"

	"github.com/w1kend/guavapay_test/internal/pkg/api/parcel_delivery_grpc"
	"github.com/w1kend/guavapay_test/internal/repositories"
)

type Implementation struct {
	parcel_delivery_grpc.UnimplementedParcelDeliveryServer
	OrdersRepo repositories.Orders
}

func NewImplementation(db *sql.DB) Implementation {
	return Implementation{
		OrdersRepo: repositories.NewOrdersRepo(db),
	}
}

// var _ parcel_delivery.ParcelDeliveryServer = &Implementation{}
