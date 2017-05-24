package store

import (
	"context"
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
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, s Store) {
	c.Set(key, s)
}

func NewContext(ctx context.Context, s Store) context.Context {
	return context.WithValue(ctx, key, s)
}
