package handlers

import (
	"net/http"

	"github.com/pressly/chi/render"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": "Ok",
	})

	return nil
}
