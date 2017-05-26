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

	// run query with obj values
	// then override obj with returned value
	err = stmt.Get(obj, obj)
	if err != nil {
		return errors.Wrap(err, "sqlx: cannot execute query")
	}

	return nil
}
