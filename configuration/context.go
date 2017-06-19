package configuration

import (
	"context"
)

const key = "configuration"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) Configuration {
	return c.Value(key).(Configuration)
}

func NewContext(ctx context.Context, cfg Configuration) context.Context {
	return context.WithValue(ctx, key, cfg)
}
