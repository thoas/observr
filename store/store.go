package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/thoas/observr/configuration"
)

type DataStore struct {
	connection sqalx.Node
}

func NewDataStore(cfg configuration.Data) (*DataStore, error) {
	dbx, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to postgres server")
	}

	dbx.SetMaxIdleConns(cfg.MaxIdleConnections)
	dbx.SetMaxOpenConns(cfg.MaxOpenConnections)

	node, err := sqalx.New(dbx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot instantiate postgres client driver")
	}

	return &DataStore{
		connection: node,
	}, nil
}

func Load(cfg configuration.Data) (*DataStore, error) {
	return NewDataStore(cfg)
}

// Connection returns SQLStore current connection.
func (s *DataStore) Connection() sqalx.Node {
	return s.connection
}

// Ping pings the storage to know if it's alive.
func (s *DataStore) Ping() error {
	row, err := s.Connection().Query("SELECT true")
	if row != nil {
		defer func() {
			// Cannot captures or logs this error.
			thr := row.Close()
			_ = thr
		}()
	}
	if err != nil {
		return errors.Wrap(err, "cannot ping database")
	}
	return nil
}

// IsErrNoRows returns if given error is a "no rows" error.
func IsErrNoRows(err error) bool {
	return errors.Cause(err) == sql.ErrNoRows
}
