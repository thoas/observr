package application

import (
	"github.com/thoas/observr/config"
	"github.com/thoas/observr/events"
	"github.com/thoas/observr/logger"
	"github.com/thoas/observr/store"
	"golang.org/x/net/context"
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

	eventStore, err := events.Load(cfg.Events)

	if err != nil {
		return nil, err
	}

	ctx := store.NewContext(context.Background(), *dataStore)
	ctx = events.NewContext(ctx, *eventStore)
	ctx = logger.NewContext(ctx, *log)
	ctx = config.NewContext(ctx, *cfg)

	return ctx, nil
}
