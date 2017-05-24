package worker

import (
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"go.uber.org/zap"

	funk "github.com/thoas/go-funk"
	"github.com/thoas/observr/broker"
	"github.com/thoas/observr/logger"

	"golang.org/x/net/context"
)

var handlers = map[string]broker.Handler{
	"observr.user.created": ErrorHandler(UserCreatedHandler),
}

func Run(ctx context.Context) error {
	var (
		wg = &sync.WaitGroup{}
		c  = make(chan os.Signal, 1)
		l  = logger.FromContext(ctx)
		b  = broker.FromContext(ctx)
	)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	wg.Add(1)

	keys := funk.Keys(handlers).([]string)

	l.Info("Starting worker", zap.String("tasks", strings.Join(keys, ", ")))

	err := b.Run(ctx, handlers)

	if err != nil {
		return err
	}

	go func() {
		for s := range c {
			l.Info("Gracefully stopping the worker", zap.Stringer("signal", s))

			b.Stop()
			wg.Done()
		}
	}()

	wg.Wait()

	return nil
}
