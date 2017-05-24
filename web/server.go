package web

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thoas/observr/configuration"
)

func Run(ctx context.Context) error {
	cfg := configuration.FromContext(ctx)

	r, err := Routes(ctx)

	if err != nil {
		return err
	}

	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
