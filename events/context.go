package events

import (
	"golang.org/x/net/context"
)

const key = "events"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) EventStore {
	return c.Value(key).(EventStore)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, e EventStore) {
	c.Set(key, e)
}

func NewContext(ctx context.Context, e EventStore) context.Context {
	return context.WithValue(ctx, key, e)
}
