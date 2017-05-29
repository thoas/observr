package web

import (
	"context"
	"fmt"

	"github.com/thoas/observr/configuration"
)

func Run(ctx context.Context) error {
	cfg := configuration.FromContext(ctx)

	r, err := Routes(ctx)

	if err != nil {
		return err
	}

	return r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
