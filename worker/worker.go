package worker

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/thoas/observr/broker"

	"golang.org/x/net/context"
)

func Run(ctx context.Context) {
	var (
		wg = &sync.WaitGroup{}
		c  = make(chan os.Signal, 1)
	)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	wg.Add(1)

	b := broker.FromContext(ctx)
	b.Run(ctx, map[string]broker.Handler{
		"observr.user.created": ErrorHandler(UserCreatedHandler),
	})

	go func() {
		for _ = range c {
			b.Stop()
			wg.Done()
		}
	}()

	wg.Wait()
}
