package events

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
)

type EventStore struct {
	Producer *nsq.Producer
}

func NewEventStore(addr string, config *nsq.Config) (*EventStore, error) {
	w, err := nsq.NewProducer(addr, config)

	if err != nil {
		return nil, err
	}

	err = w.Ping()

	if err != nil {
		return nil, err
	}

	return &EventStore{
		Producer: w,
	}, nil
}

func (ns *EventStore) Publish(event Event) error {
	msg, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ns.Producer.Publish(event.Name(), msg)
}
