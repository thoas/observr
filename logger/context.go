package logger

import (
	"context"

	"go.uber.org/zap"
)

const key = "logger"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) zap.Logger {
	return c.Value(key).(zap.Logger)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, l zap.Logger) {
	c.Set(key, l)
}

func NewContext(ctx context.Context, l zap.Logger) context.Context {
	return context.WithValue(ctx, key, l)
}
