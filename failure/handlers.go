package failure

import (
	"net/http"

	"github.com/mholt/binding"
	"github.com/pkg/errors"
	"github.com/pressly/chi/render"
	"github.com/thoas/observr/store"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func HandleError(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)

		if err != nil {
			ProcessError(w, r, err)
			return
		}
	}
}

func wrapError(err error) error {
	switch e := err.(type) {
	case binding.Errors:
		return ValidationError(e)
	}

	return err
}

func ProcessError(w http.ResponseWriter, r *http.Request, err error) {
	cause := wrapError(errors.Cause(err))

	switch e := cause.(type) {
	case HTTPError:
		render.Status(r, e.Status)
		render.JSON(w, r, e)
	default:
		if store.IsErrNoRows(cause) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "resource not found"})
		}
	}
}
