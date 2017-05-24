package broker

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/thoas/observr/configuration"
	"github.com/ulule/amqpx"
)

const (
	// WarpExchange transfer expired delayed messages to direct queue
	// We use default amq.direct exchange for convenience
	WarpExchange = "amq.direct"
)

// AMQPBroker service
type AMQPBroker struct {
	pool      amqpx.Pooler
	consumers []*AMQPConsumer
}

// NewAMQPBroker creates an AMQP broker
func NewAMQPBroker(cfg configuration.Broker) (*AMQPBroker, error) {
	pool, err := amqpx.NewChannelPool(func() (*amqp.Connection, error) {
		return amqp.Dial(cfg.URI)
	}, amqpx.Bounds(25, 50))

	if err != nil {
		return nil, errors.Wrap(err, "cannot create a amqp connection pool")
	}

	return &AMQPBroker{
		pool: pool,
	}, nil
}

func newAMQPMessage(body []byte) amqp.Publishing {
	return amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		Timestamp:       time.Now(),
		ContentType:     "application/json",
		ContentEncoding: "utf-8",
		Body:            body,
	}
}

func (broker *AMQPBroker) Run(ctx context.Context, handlers map[string]Handler) error {
	var (
		consumers = make([]*AMQPConsumer, len(handlers))
		i         = 0
	)

	for key, handler := range handlers {
		consumer, err := NewAMQPConsumer(ctx, broker.pool, key, handler)

		if err != nil {
			return err
		}

		err = consumer.Start(ctx)

		if err != nil {
			return err
		}

		consumers[i] = consumer

		i += 1
	}

	return nil
}

func getDirectChannel(connection amqpx.Connector, directTopic string) (*amqp.Channel, error) {
	directChannel, err := connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "cannot create a new direct channel")
	}

	directQueue, err := directChannel.QueueDeclare(
		directTopic,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot declare topic")
	}

	err = directChannel.QueueBind(
		directQueue.Name, // queue name
		directQueue.Name, // routing key
		WarpExchange,     // exchange
		false,            // noWait
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot bind exchange")
	}

	return directChannel, nil
}

func getDelayedChannel(connection amqpx.Connector, delayedTopic string, fallbackChannel string) (*amqp.Channel, error) {

	delayedChannel, err := connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "cannot create a new delayed channel")
	}

	_, err = delayedChannel.QueueDeclare(
		delayedTopic,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		amqp.Table{
			"x-dead-letter-exchange":    WarpExchange,
			"x-dead-letter-routing-key": fallbackChannel,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot declare a new topic")
	}

	return delayedChannel, nil
}

// Publish emits a message to a given topic
func (broker *AMQPBroker) Stop() {
	for _, consumer := range broker.consumers {
		consumer.Stop()
	}
}

// Publish emits a message to a given topic
func (broker *AMQPBroker) Publish(ctx context.Context, ev Event) error {
	connection, err := broker.pool.Get()
	if err != nil {
		return errors.Wrap(err, "cannot acquire a new AMQP connection from pool")
	}

	defer connection.Close()

	directTopic := ev.Name()

	directChannel, err := getDirectChannel(connection, directTopic)
	if err != nil {
		return errors.Wrapf(err, "cannot acquire a channel for topic: %s", directTopic)
	}

	defer directChannel.Close()

	json, err := ev.ToBytes()
	if err != nil {
		return errors.Wrapf(err, "cannot encode %T event as json", ev)
	}

	msg := newAMQPMessage(json)

	err = directChannel.Publish(
		"",          // default exchange
		directTopic, // queue name
		false,       // mandatory
		false,       // immediate
		msg,
	)
	if err != nil {
		return errors.Wrapf(err, "cannot publish on queue: %s", directTopic)
	}

	return nil
}

// AMQPConsumer listen on a given queue for AMQP events
type AMQPConsumer struct {
	Channel   *amqp.Channel
	QueueName string
	Consumer  string
	Handler   Handler
	done      chan error
}

// NewAMQPConsumer returns an AMQP Consumer
func NewAMQPConsumer(ctx context.Context, pool amqpx.Pooler, queueName string, handler Handler) (*AMQPConsumer, error) {
	connection, err := pool.Get()
	if err != nil {
		return nil, errors.Wrap(err, "cannot acquire a new connection from pool")
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "cannot acquire a channel")
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot declare topic: %s", queueName)
	}

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, errors.Wrap(err, "cannot configure channel")
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, errors.Wrap(err, "cannot obtain server hostname")
	}

	consumer := &AMQPConsumer{
		Channel:   channel,
		Handler:   handler,
		QueueName: queue.Name,
		Consumer:  fmt.Sprintf("consumer-%s#%d@%s", queue.Name, os.Getpid(), hostname),
		done:      make(chan error),
	}

	return consumer, nil
}

// Start launch events listening
func (consumer *AMQPConsumer) Start(ctx context.Context) error {
	deliveries, err := consumer.Channel.Consume(
		consumer.QueueName,
		consumer.Consumer, // uniq consumer id
		false,             // noAck
		false,             // exclusive
		false,             // noLocal
		false,             // noWait
		nil,               // arguments
	)
	if err != nil {
		return errors.Wrap(err, "cannot create a message consumer")
	}

	go func() {
		for delivery := range deliveries {
			consumer.Handler(ctx, delivery.Body)

			// only ack message if successfully processed
			delivery.Ack(false)
		}

		consumer.done <- nil
	}()

	return nil
}

// Stop interrupts events listening
func (consumer *AMQPConsumer) Stop() error {
	err := consumer.Channel.Cancel(consumer.Consumer, true)
	if err != nil {
		return errors.Wrap(err, "cannot shutdown consumer")
	}

	return <-consumer.done
}
