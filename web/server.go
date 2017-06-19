package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thoas/observr/configuration"
)

func Run(ctx context.Context) error {
	cfg := configuration.FromContext(ctx)

	r, err := Routes(ctx)

	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), r)
}
