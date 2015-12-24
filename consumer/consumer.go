package consumer

import (
	"github.com/Sirupsen/logrus"
	"github.com/nsqio/go-nsq"
	"github.com/thoas/observr/worker"
)

type Consumer struct {
	consumers []*nsq.Consumer
	worker    *worker.Worker
}

type Option struct {
	Topic    string
	Channel  string
	Addr     string
	Config   *nsq.Config
	Logger   *logrus.Logger
	Handlers []nsq.HandlerFunc
}

func NewConsumer(option *Option) (*Consumer, error) {
	consumer, err := nsq.NewConsumer(option.Topic, option.Channel, option.Config)

	if err != nil {
		return nil, err
	}

	for _, f := range option.Handlers {
		consumer.AddHandler(f)
	}

	worker, err := worker.New(option.Addr, option.Logger, consumer)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		worker:    worker,
		consumers: []*nsq.Consumer{consumer},
	}, nil
}

func (c *Consumer) Consume() {
	c.worker.Run()
}
