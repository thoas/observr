package web

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/web/handlers"
)

func Routes(ctx context.Context) (*gin.Engine, error) {
	r := gin.Default()

	r.GET("/healthcheck", failure.HandleError(handlers.HealthcheckHandler))

	return r, nil
}
