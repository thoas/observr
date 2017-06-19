package store

import (
	"context"

	"github.com/heetch/sqalx"
)

const key = "data"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

type Store interface {
	Close() error
	Ping() error
	Flush() error
	Connection() sqalx.Node
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

func NewContext(ctx context.Context, s Store) context.Context {
	return context.WithValue(ctx, key, s)
}
