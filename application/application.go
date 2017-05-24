package application

import (
	"context"

	"github.com/thoas/observr/broker"
	"github.com/thoas/observr/config"
	"github.com/thoas/observr/logger"
	"github.com/thoas/observr/store"
)

func Load(path string) (context.Context, error) {
	cfg, err := config.Load(path)

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

	ctx := store.NewContext(context.Background(), *dataStore)
	ctx = broker.NewContext(ctx, b)
	ctx = logger.NewContext(ctx, *log)
	ctx = config.NewContext(ctx, *cfg)

	return ctx, nil
}
