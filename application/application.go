package application

import (
	"github.com/Sirupsen/logrus"
	"github.com/nsqio/go-nsq"
	"github.com/thoas/observr/consumer"
	"github.com/thoas/observr/events"
	"github.com/thoas/observr/store"
)

type Application struct {
	DataStore  *store.DataStore
	EventStore *events.EventStore
	Consumer   *consumer.Consumer
	Logger     *logrus.Logger
}

func New() (*Application, error) {
	option := &store.Option{
		Name: "observ",
		Ips:  []string{"127.0.0.1"},
	}

	dataStore, err := store.NewDataStore(option)

	if err != nil {
		return nil, err
	}

	logger := logrus.New()

	app := &Application{
		DataStore: dataStore,
		Logger:    logger,
	}

	config := nsq.NewConfig()

	tcpAddr := "127.0.0.1:4150"

	httpAddr := "127.0.0.1:4161"

	consumer, err := consumer.NewConsumer(&consumer.Option{
		HTTPAddrs: []string{httpAddr},
		Topic:     "test",
		Channel:   "observr",
		Logger:    logger,
		Config:    config,
		Handlers: []nsq.HandlerFunc{
			app.Handle(TestHandler),
		},
	})

	if err != nil {
		return nil, err
	}

	app.Consumer = consumer

	eventStore, err := events.NewEventStore(tcpAddr, config)

	if err != nil {
		return nil, err
	}

	app.EventStore = eventStore

	return app, nil
}

func (a *Application) Handle(h Handler) nsq.HandlerFunc {
	return nsq.HandlerFunc(func(message *nsq.Message) error {
		return h(a, message)
	})
}

func (a *Application) Work() {
	a.Consumer.Consume()
}
