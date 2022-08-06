package config

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	"github.com/pkg/errors"
	"github.com/w1kend/parcel_delivery/internal/pkg/health"
)

func NewPostgres() (*sql.DB, error) {
	conn, err := sql.Open("postgres", GetValue(PgDsn).String())
	if err != nil {
		return nil, errors.Wrap(err, "conn to postgres")
	}

	if conn.Ping() != nil {
		return nil, errors.Wrap(err, "ping postgres")
	}

	conn.SetMaxOpenConns(GetValue(PgMaxOpenConns).Int())
	conn.SetMaxIdleConns(GetValue(PgMaxIddleConns).Int())

	return conn, nil
}

var (
	pgOnce  sync.Once
	testPg  sql.DB
	testMux sync.RWMutex
)

// NewTestPostgres - connects to postgress for tests
func NewTestPostgres(t *testing.T) *sql.DB {
	pgOnce.Do(func() {
		testMux.Lock()
		defer testMux.Unlock()

		conn, err := sql.Open("postgres", GetValue(PgDsn).String())
		if err != nil {
			t.Fatalf("failed to connect to postgres: %s", err.Error())
		}

		testPg = *conn
	})

	testMux.RLock()
	defer testMux.RUnlock()
	return &testPg
}

func PgHealthChecker(conn *sql.DB) health.CheckFn {
	return func() health.State {
		if conn == nil {
			return health.State{
				Status:      health.StatusWarn,
				Description: "can't check postgress connection",
			}
		}
		if err := conn.Ping(); err != nil {
			return health.State{
				Status:      health.StatusCrit,
				Description: fmt.Sprintf("ping to postgress: %s", err.Error()),
			}
		}

		return health.State{
			Status:      health.StatusOK,
			Description: "OK",
		}
	}
}
