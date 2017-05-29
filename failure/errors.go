package failure

import (
	"net/http"

	"github.com/mholt/binding"
)

type HTTPError struct {
	Status  int            `json:"-"`
	Message interface{}    `json:"message"`
	Type    string         `json:"type,omitempty"`
	Err     error          `json:"-"`
	Errors  binding.Errors `json:"errors,omitempty"`
}

func (e HTTPError) Error() string {
	return e.Message.(string)
}

func ValidationError(errs binding.Errors) error {
	code := 422

	if len(errs) > 0 && errs[0].Kind() == "ContentTypeError" {
		code = http.StatusUnsupportedMediaType
	}

	return HTTPError{
		Status:  code,
		Message: ValidationFailedMessage,
		Errors:  errs,
		Type:    "validation_failed",
	}
}

// AlreadyExists on duplicate resources.
func AlreadyExists(fields []string) error {
	errs := binding.Errors{}
	errs.Add(fields, "AlreadyExistsError", AlreadyExistsMessage)

	return ValidationError(errs)
}

func NotFoundError() error {
	return HTTPError{
		Status:  http.StatusNotFound,
		Message: NotFoundErrorMessage,
		Type:    "not_found",
	}
}

func PermissionError() error {
	return HTTPError{
		Status:  http.StatusUnauthorized,
		Message: PermissionDeniedMessage,
		Type:    "permission_denied",
	}
}
