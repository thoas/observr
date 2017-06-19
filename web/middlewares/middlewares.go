package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/thoas/observr/configuration"
	"github.com/thoas/observr/store"
)

func Application(appctx context.Context) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			ctx = configuration.NewContext(ctx, configuration.FromContext(appctx))
			ctx = store.NewContext(ctx, store.FromContext(appctx))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}

func APIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		ctx := r.Context()

		if header != "" {
			parts := strings.Split(header, " ")

			if len(parts) == 2 && strings.ToLower(parts[0]) == "apikey" {
				user, err := store.GetUserByAPIKey(ctx, parts[1])

				if err != nil {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}

				ctx = context.WithValue(ctx, "user", user)
			} else {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
