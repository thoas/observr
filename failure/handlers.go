package failure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
	"github.com/pkg/errors"
	"github.com/thoas/observr/store"
)

func HandleError(handler func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			ProcessError(c, err)
			c.Abort()
			return
		}

		c.Next()
	}
}

func wrapError(err error) error {
	switch e := err.(type) {
	case binding.Errors:
		return ValidationError(e)
	}

	return err
}

func ProcessError(c *gin.Context, err error) {
	cause := wrapError(errors.Cause(err))

	switch e := cause.(type) {
	case HTTPError:
		c.JSON(e.Status, e)
	default:
		if store.IsErrNoRows(cause) {
			c.JSON(http.StatusNotFound, "resource not found")
		}
	}
}
