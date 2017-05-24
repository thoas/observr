package failure

import (
	"net/http"

	"github.com/mholt/binding"
)

type HttpError struct {
	Status  int            `json:"-"`
	Message interface{}    `json:"message"`
	Type    string         `json:"type,omitempty"`
	Err     error          `json:"-"`
	Errors  binding.Errors `json:"errors,omitempty"`
}

func (e HttpError) Error() string {
	return e.Message.(string)
}

func ValidationError(errs binding.Errors) error {
	code := 422

	if len(errs) > 0 && errs[0].Kind() == "ContentTypeError" {
		code = http.StatusUnsupportedMediaType
	}

	return HttpError{
		Status:  code,
		Message: ValidationFailed,
		Errors:  errs,
		Type:    "validation_failed",
	}
}
