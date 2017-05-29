package web

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thoas/observr/configuration"
	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/web/handlers"
	"github.com/thoas/observr/web/middlewares"
)

func Routes(ctx context.Context) (*gin.Engine, error) {
	cfg := configuration.FromContext(ctx)

	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middlewares.Application(ctx))

	r.GET("/healthcheck", failure.HandleError(handlers.Healthcheck))
	r.POST("/users", failure.HandleError(handlers.UserCreate))

	return r, nil
}
