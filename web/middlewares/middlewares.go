package middlewares

import (
	"context"
	"net/http"
	"strings"

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

func APIKey(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header != "" {
			parts := strings.Split(header, " ")

			if len(parts) == 2 && strings.ToLower(parts[0]) == "apikey" {
				user, err := store.GetUserByAPIKey(c, parts[1])

				if err != nil {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}

				c.Set("user", user)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		c.Next()
	}
}
