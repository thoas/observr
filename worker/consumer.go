package worker

import (
	"github.com/nsqio/go-nsq"
	"golang.org/x/net/context"
)

type Consumer struct {
	*nsq.Consumer
	handler Handler
	ctx     context.Context
}

func NewConsumer(topic string, channel string, config *nsq.Config, ctx context.Context, handler Handler) (*Consumer, error) {
	consumer, err := nsq.NewConsumer(topic, channel, config)

	if err != nil {
		return nil, err
	}

	c := &Consumer{
		consumer,
		handler,
		ctx,
	}

	consumer.AddHandler(c.Handle())

	return c, nil
}

func (c *Consumer) Handle() nsq.HandlerFunc {
	return nsq.HandlerFunc(func(message *nsq.Message) error {
		return c.handler(message, c.ctx)
	})
}
