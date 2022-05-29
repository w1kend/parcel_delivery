package repositories

import (
	"context"
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/table"
)

type Orders interface {
	CreateOrder(ctx context.Context, order model.Orders) error
	GetOrder(ctx context.Context, id string) (*model.Orders, error)
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
	stmt := table.Orders.INSERT(table.Orders.AllColumns).MODEL(order)

	query, args := stmt.Sql()
	_, err := r.db.QueryContext(ctx, query, args...)

	return err
}

// GetOrder returns order by id
func (r DBOrders) GetOrder(ctx context.Context, id string) (*model.Orders, error) {
	stmt := table.Orders.
		SELECT(table.Orders.AllColumns).
		WHERE(table.Orders.ID.EQ(postgres.StringExp(postgres.CAST(postgres.String(id)).AS("uuid"))))

	var order model.Orders
	err := stmt.QueryContext(ctx, r.db, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

type OrdersRepoFilter struct {
	CreatedBy *int64
}
