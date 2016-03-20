package events

import (
	"github.com/nsqio/go-nsq"
	"github.com/thoas/observr/config"
)

type Event interface {
	Name() string
}

func Load(cfg config.Events) (*EventStore, error) {
	config := nsq.NewConfig()

	return NewEventStore(cfg.Producer, config)
}
