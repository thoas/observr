package store

import (
	"golang.org/x/net/context"
)

const key = "data"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) DataStore {
	return c.Value(key).(DataStore)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, s DataStore) {
	c.Set(key, s)
}

func NewContext(ctx context.Context, s DataStore) context.Context {
	return context.WithValue(ctx, key, s)
}
