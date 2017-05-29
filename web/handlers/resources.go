package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/store"
)

func UserResource() gin.HandlerFunc {
	return failure.HandleError(func(c *gin.Context) error {
		id := c.Param("id")
		if id == "" {
			return errors.Wrap(failure.NotFoundError(), "cannot find user")
		}

		user, err := store.GetUserByID(c, id)
		if err != nil {
			return err
		}

		c.Set("resource", user)

		return nil
	})
}
