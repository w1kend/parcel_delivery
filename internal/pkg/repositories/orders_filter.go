package repositories

import (
	"github.com/go-jet/jet/v2/postgres"
	"github.com/pkg/errors"
	"github.com/w1kend/parcel_delivery/internal/generated/parcel_delivery/public/table"
)

// OrdersFilter - filter for orders repository
type OrdersFilter struct {
	OrderID   *string
	CreatedBy *string
	CourierID *string
	//offset
	//limit
}

func (f OrdersFilter) build() (postgres.BoolExpression, error) {
	if err := f.validate(); err != nil {
		return nil, err
	}

	var whereStmt postgres.BoolExpression
	if f.OrderID != nil {
		whereStmt = whereStmt.AND(table.Orders.ID.EQ(uuidExpr(*f.OrderID)))
	}

	// if filter.CourierID != nil {
	// 	whereStmt = whereStmt.AND(table.Orders.CourierID.EQ(uuidExpr(*filter.CourierID)))
	// }

	if f.CreatedBy != nil {
		whereStmt = whereStmt.AND(table.Orders.CreatedBy.EQ(uuidExpr(*f.CreatedBy)))
	}

	return whereStmt, nil
}

func (f OrdersFilter) validate() error {
	if f.OrderID == nil && f.CreatedBy == nil && f.CourierID == nil {
		return errors.New("at least one parameter is required")
	}

	return nil
}
