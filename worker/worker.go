package worker

import (
	"github.com/Sirupsen/logrus"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	wg        *sync.WaitGroup
	consumers []*nsq.Consumer
	logger    *logrus.Logger
	timeout   time.Duration
}

func New(tcpAddrs []string, httpAddrs []string, logger *logrus.Logger, consumers ...*nsq.Consumer) (*Worker, error) {
	wg := &sync.WaitGroup{}

	for _, consumer := range consumers {
		err := consumer.ConnectToNSQDs(tcpAddrs)

		if err != nil {
			return nil, err
		}

		err = consumer.ConnectToNSQLookupds(httpAddrs)

		if err != nil {
			return nil, err
		}

		wg.Add(1)
	}

	return &Worker{
		wg,
		consumers,
		logger,
		1 * time.Minute,
	}, nil
}

func (w *Worker) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for {
			sig := <-c
			w.logger.WithFields(logrus.Fields{
				"signal": sig,
			}).Info("receive signal")

			for _, c := range w.consumers {
				w.logger.Info("stopping consumer... ")
				c.Stop()

				select {
				case <-c.StopChan:
					w.logger.Info("consumer stopped")
					w.wg.Done()
				case <-time.After(w.timeout):
					w.logger.WithFields(logrus.Fields{
						"timeout": w.timeout,
					}).Warn("timeout while stopping consumer")
				}
			}
		}
	}()

	w.wg.Wait()
}
