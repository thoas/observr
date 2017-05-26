package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thoas/observr/configuration"
	"github.com/thoas/observr/store"
)

func Application(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		configuration.ToContext(c, configuration.FromContext(ctx))
		store.ToContext(c, store.FromContext(ctx))

		c.Next()
	}
}
