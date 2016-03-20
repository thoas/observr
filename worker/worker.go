package worker

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/Sirupsen/logrus"
	"github.com/nsqio/go-nsq"
	"github.com/thoas/observr/application"
	"github.com/thoas/observr/config"
	"github.com/thoas/observr/logger"
)

type Worker struct {
	wg        *sync.WaitGroup
	consumers []*Consumer
	ctx       context.Context
	timeout   time.Duration
}

func Load(path string) (*Worker, error) {
	ctx, err := application.Load(path)

	if err != nil {
		return nil, err
	}

	return New(ctx)
}

func New(ctx context.Context) (*Worker, error) {
	wg := &sync.WaitGroup{}

	cfg := config.FromContext(ctx)

	nsqConfig := nsq.NewConfig()

	tasks := map[string]Handler{
		"test": TestHandler,
	}

	var consumers []*Consumer

	for name, handler := range tasks {
		consumer, err := NewConsumer(name, "tasks", nsqConfig, ctx, handler)

		if err != nil {
			return nil, err
		}

		err = consumer.ConnectToNSQLookupds(cfg.Events.Consumer)

		if err != nil {
			return nil, err
		}

		consumers = append(consumers, consumer)

		wg.Add(1)
	}

	return &Worker{
		wg,
		consumers,
		ctx,
		1 * time.Minute,
	}, nil
}

func (w *Worker) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for {
			sig := <-c

			log := logger.FromContext(w.ctx)

			log.WithFields(logrus.Fields{
				"signal": sig,
			}).Info("receive signal")

			for _, c := range w.consumers {
				log.Info("stopping consumer... ")
				c.Stop()

				select {
				case <-c.StopChan:
					log.Info("consumer stopped")
					w.wg.Done()
				case <-time.After(w.timeout):
					log.WithFields(logrus.Fields{
						"timeout": w.timeout,
					}).Warn("timeout while stopping consumer")
				}
			}
		}
	}()

	w.wg.Wait()
}
