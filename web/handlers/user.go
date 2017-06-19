package handlers

import (
	"net/http"

	"github.com/mholt/binding"
	"github.com/pressly/chi/render"

	"github.com/thoas/observr/manager"
	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/rpc/resources"
	"github.com/thoas/observr/store/models"
)

func UserCreate(w http.ResponseWriter, r *http.Request) error {
	var payload payloads.UserCreate

	errs := binding.Bind(r, &payload)
	if errs != nil {
		return errs
	}

	ctx := r.Context()

	user, err := manager.CreateUser(ctx, &payload)
	if err != nil {
		return err
	}

	resource, err := resources.NewUser(ctx, user)
	if err != nil {
		return err
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resource)

	return nil
}

func ProjectCreate(w http.ResponseWriter, r *http.Request) error {
	var payload payloads.ProjectCreate

	errs := binding.Bind(r, &payload)

	if errs != nil {
		return errs
	}

	ctx := r.Context()

	user := ctx.Value("user").(*models.User)

	project, err := manager.CreateProject(ctx, &payload, user)
	if err != nil {
		return err
	}

	resource, err := resources.NewProject(ctx, project)
	if err != nil {
		return err
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resource)

	return nil
}
