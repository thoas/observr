package config

import (
	"golang.org/x/net/context"
)

const key = "config"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) Config {
	return c.Value(key).(Config)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, cfg Config) {
	c.Set(key, cfg)
}

func NewContext(ctx context.Context, cfg Config) context.Context {
	return context.WithValue(ctx, key, cfg)
}
