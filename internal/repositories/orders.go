package repositories

import (
	"context"
	"database/sql"

	"github.com/w1kend/guavapay_test/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/guavapay_test/internal/generated/parcel_delivery/public/table"
)

type Orders interface {
	CreateOrder(ctx context.Context, order model.Orders) error
}

type DBOrders struct {
	db *sql.DB
}

// NewOrdersRepo returns Orders implementation
func NewOrdersRepo(db *sql.DB) DBOrders {
	return DBOrders{db: db}
}

var _ Orders = DBOrders{}

// CreateOrder implements Orders
func (r DBOrders) CreateOrder(ctx context.Context, order model.Orders) error {
	stmt := table.Orders.INSERT(table.Orders.MutableColumns).MODEL(order)

	_, err := r.db.Exec(stmt.Sql())
	return err
}
