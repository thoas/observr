package failure

import (
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
	"github.com/pkg/errors"
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

	return nil
}

func ProcessError(c *gin.Context, err error) {
	// add error in the Context (automatic log)
	c.Error(err)

	cause := wrapError(errors.Cause(err))

	switch e := cause.(type) {
	case HttpError:
		c.JSON(e.Status, e)
	}
}
