package logger

import "github.com/Sirupsen/logrus"

func Load() *logrus.Logger {
	return logrus.New()
}
