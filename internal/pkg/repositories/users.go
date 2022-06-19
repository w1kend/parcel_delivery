package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/qrm"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/model"
	"github.com/w1kend/parcel_delivery_test/internal/generated/parcel_delivery/public/table"
)

// Users repository interface
type Users interface {
	CreateUser(ctx context.Context, user model.Users) error
	GetByEmail(ctx context.Context, email string) (*model.Users, error)
}

// DBUsers implements Users repository
type DBUsers struct {
	db *sql.DB
}

var _ Users = DBUsers{}

// NewUsersRepository returns a new instance of UsersRepository
func NewUsersRepository(db *sql.DB) DBUsers {
	return DBUsers{
		db: db,
	}
}

// CreateUser - creates user
func (d DBUsers) CreateUser(ctx context.Context, user model.Users) error {
	stmt := table.Users.
		INSERT(
			table.Users.ID,
			table.Users.Name,
			table.Users.Email,
			table.Users.Role,
			table.Users.Password,
			table.Users.CreatedAt,
		).MODEL(user)

	_, err := stmt.ExecContext(ctx, d.db)
	return err
}

// GetByEmail - returns user with specified email
func (d DBUsers) GetByEmail(ctx context.Context, email string) (*model.Users, error) {
	stmt := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.Email.EQ(postgres.String(email))).
		LIMIT(1)

	var user model.Users
	err := stmt.QueryContext(ctx, d.db, &user)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
