package handlers_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
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
