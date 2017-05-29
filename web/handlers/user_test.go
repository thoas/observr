package handlers_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thoas/observr/store"
	"github.com/thoas/observr/store/models"
	"github.com/thoas/observr/test"
)

func TestHandlers_User_UserCreateHandler(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		resp := test.GET(ctx, "/users")

		is.Equal(http.StatusNotFound, resp.Code)

		resp = test.POST(ctx, "/users", nil)

		is.Equal(http.StatusUnsupportedMediaType, resp.Code)

		resp = test.POST(ctx, "/users", map[string]interface{}{
			"username": "foo",
			"email":    "foo@bar.com",
			"password": "$ecret",
		})

		is.Equal(http.StatusCreated, resp.Code)
	})
}

func TestHandlers_User_ProjectCreateHandler(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		resp := test.POST(ctx, fmt.Sprintf("/users/%s/projects", store.UUID()), map[string]interface{}{
			"name": "Ulule",
		})

		is.Equal(http.StatusNotFound, resp.Code)

		u := &models.User{
			Username: "thoas",
			Email:    "florent@ulule.com",
			Password: "$ecret",
		}

		err := store.CreateUser(ctx, u)
		is.Nil(err)

		resp = test.POST(ctx, fmt.Sprintf("/users/%s/projects", u.ID), map[string]interface{}{
			"name": "Ulule",
		})

		is.Equal(http.StatusCreated, resp.Code)
	})
}
