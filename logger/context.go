package logger

import (
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

const key = "logger"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) logrus.Logger {
	return c.Value(key).(logrus.Logger)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, l logrus.Logger) {
	c.Set(key, l)
}

func NewContext(ctx context.Context, l logrus.Logger) context.Context {
	return context.WithValue(ctx, key, l)
}
