package store

import (
	"context"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	return uuid.NewV4().String()
}

func Sync(ctx context.Context, query string, obj interface{}) error {
	strx := FromContext(ctx)

	stmt, err := strx.Connection().PrepareNamed(query)
	if err != nil {
		return errors.Wrap(err, "sqlx: cannot prepare statement")
	}
	defer stmt.Close()

	err = stmt.Get(obj, obj)
	if err != nil {
		return errors.Wrap(err, "sqlx: cannot execute query")
	}

	return nil
}

func GetByParams(ctx context.Context, dest interface{}, query string, params map[string]interface{}) error {
	strx := FromContext(ctx)

	stmt, err := strx.Connection().PrepareNamed(query)
	if err != nil {
		return errors.Wrap(err, "sqlx: cannot prepare statement")
	}

	err = stmt.Get(dest, params)
	if err != nil {
		return errors.Wrap(err, "sqlx: cannot execute query")
	}

	return nil
}

func GetByID(ctx context.Context, dest interface{}, query string, id string) error {
	args := map[string]interface{}{
		"id": id,
	}

	return GetByParams(ctx, dest, query, args)
}
