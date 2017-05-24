package application

import (
	"context"

	"github.com/thoas/observr/broker"
	"github.com/thoas/observr/configuration"
	"github.com/thoas/observr/logger"
	"github.com/thoas/observr/store"
)

func Load(path string) (context.Context, error) {
	cfg, err := configuration.Load(path)

	if err != nil {
		return nil, err
	}

	dataStore, err := store.Load(cfg.Data)

	if err != nil {
		return nil, err
	}

	log := logger.Load()

	b, err := broker.Load(cfg.Broker)

	if err != nil {
		return nil, err
	}

	ctx := store.NewContext(context.Background(), dataStore)
	ctx = broker.NewContext(ctx, b)
	ctx = logger.NewContext(ctx, *log)
	ctx = configuration.NewContext(ctx, *cfg)

	return ctx, nil
}

func Shutdown(ctx context.Context) error {
	s := store.FromContext(ctx)

	err := s.Flush()
	if err != nil {
		return err
	}

	return s.Close()
}
