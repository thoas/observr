package web

import (
	"context"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/web/handlers"
	"github.com/thoas/observr/web/middlewares"
)

func Routes(ctx context.Context) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(middlewares.Application(ctx))
	r.Use(middlewares.APIKey)

	r.Get(
		"/healthcheck",
		failure.HandleError(handlers.Healthcheck))
	r.Post(
		"/users",
		failure.HandleError(handlers.UserCreate))

	r.Route("/users/:id", func(r chi.Router) {
		r.Use(handlers.RequireAuth, handlers.UserResource)
		r.Post(
			"/projects",
			failure.HandleError(handlers.ProjectCreate))
	})

	return r, nil
}
