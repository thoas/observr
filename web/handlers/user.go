package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
	"github.com/thoas/observr/manager"
	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/rpc/resources"
)

func UserCreate(c *gin.Context) error {
	var payload payloads.UserCreate

	errs := binding.Bind(c.Request, &payload)
	if errs != nil {
		return errs
	}

	user, err := manager.CreateUser(c, &payload)
	if err != nil {
		return err
	}

	resource, err := resources.NewUser(c, user)
	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, resource)

	return nil
}
