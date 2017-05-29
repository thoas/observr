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

		resp := test.GET(ctx, &test.Request{URL: "/users"})

		is.Equal(http.StatusNotFound, resp.Code)

		resp = test.POST(ctx, &test.Request{URL: "/users"})

		is.Equal(http.StatusUnsupportedMediaType, resp.Code)

		resp = test.POST(ctx, &test.Request{URL: "/users", Data: map[string]interface{}{
			"username": "foo",
			"email":    "foo@bar.com",
			"password": "$ecret",
		}})

		is.Equal(http.StatusCreated, resp.Code)
	})
}

func TestHandlers_User_ProjectCreateHandler(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		resp := test.POST(ctx, &test.Request{URL: fmt.Sprintf("/users/%s/projects", store.UUID()), Data: map[string]interface{}{
			"name": "Ulule",
		}})

		is.Equal(http.StatusNotFound, resp.Code)

		u := &models.User{
			Username: "thoas",
			Email:    "florent@ulule.com",
			Password: "$ecret",
		}

		err := store.CreateUser(ctx, u)
		is.Nil(err)

		resp = test.POST(ctx, &test.Request{URL: fmt.Sprintf("/users/%s/projects", u.ID), Data: map[string]interface{}{
			"name": "Ulule",
			"url":  "https://www.ulule.com",
		}})

		is.Equal(http.StatusUnauthorized, resp.Code)

		resp = test.POST(ctx, &test.Request{URL: fmt.Sprintf("/users/%s/projects", u.ID), Data: map[string]interface{}{
			"name": "Ulule",
			"url":  "https://www.ulule.com",
		}, User: u})

		is.Equal(http.StatusCreated, resp.Code)
	})
}
