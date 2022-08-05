package repositories

import (
	"context"
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/pkg/errors"
	"github.com/w1kend/parcel_delivery/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery/internal/generated/parcel_delivery/public/table"
)

type Orders interface {
	CreateOrder(ctx context.Context, order model.Orders) error
	GetOrder(ctx context.Context, filter OrdersFilter) (*model.Orders, error)
	ListOrders(ctx context.Context, filter OrdersFilter) ([]model.Orders, error)
	ChangeDestination(ctx context.Context, id, newDestination string) error
	UpdateStatus(ctx context.Context, id string, newStatus model.OrderStatus) error
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

// GetOrder - returns order
func (r DBOrders) GetOrder(ctx context.Context, filter OrdersFilter) (*model.Orders, error) {
	where, err := filter.build()
	if err != nil {
		return nil, errors.Wrap(err, "build filter")
	}

	stmt := table.Orders.
		SELECT(table.Orders.AllColumns).
		WHERE(where)

	var order model.Orders
	err = stmt.QueryContext(ctx, r.db, &order)
	switch {
	case errors.Is(err, qrm.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return &order, nil
}

// ListOrders - returns list orders
func (r DBOrders) ListOrders(ctx context.Context, filter OrdersFilter) ([]model.Orders, error) {
	where, err := filter.build()
	if err != nil {
		return nil, errors.Wrap(err, "build filter")
	}

	stmt := table.Orders.
		SELECT(table.Orders.AllColumns).
		FROM(table.Orders).
		WHERE(where)

	var orders []model.Orders
	err = stmt.QueryContext(ctx, r.db, &orders)
	if errors.Is(err, qrm.ErrNoRows) {
		return orders, nil
	}

	return orders, err
}

// ChangeDestination - changes order destination
func (r DBOrders) ChangeDestination(ctx context.Context, id, newDestination string) error {
	stmt := table.Orders.
		UPDATE(table.Orders.ToAddr).
		SET(postgres.String(newDestination)).
		WHERE(table.Orders.ID.EQ(uuidExpr(id)))

	_, err := stmt.ExecContext(ctx, r.db)

	return err
}

// UpdateStatus - updates order status
func (r DBOrders) UpdateStatus(ctx context.Context, id string, newStatus model.OrderStatus) error {
	stmt := table.Orders.
		UPDATE(table.Orders.Status).
		SET(postgres.String(newStatus.String())).
		WHERE(table.Orders.ID.EQ(postgres.String(id)))

	_, err := stmt.ExecContext(ctx, r.db)
	return err
}

func uuidExpr(id string) postgres.StringExpression {
	return postgres.StringExp(postgres.CAST(postgres.String(id)).AS("uuid"))
}
