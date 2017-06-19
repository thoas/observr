package handlers

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/pressly/chi"

	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/store"
)

func UserResource(next http.Handler) http.Handler {
	return failure.HandleError(func(w http.ResponseWriter, r *http.Request) error {
		id := chi.URLParam(r, "id")
		if id == "" {
			return errors.Wrap(failure.NotFoundError(), "cannot find user")
		}

		ctx := r.Context()

		user, err := store.GetUserByID(ctx, id)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, "resource", user)
		next.ServeHTTP(w, r.WithContext(ctx))

		return nil
	})
}

func RequireAuth(next http.Handler) http.Handler {
	return failure.HandleError(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		user := ctx.Value("user")

		if user == nil {
			return errors.Wrap(failure.PermissionError(), "user not authenticated")
		}

		next.ServeHTTP(w, r.WithContext(ctx))

		return nil
	})
}
