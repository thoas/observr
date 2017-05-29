package handlers_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thoas/observr/test"
)

func TestHandlers_Base_HealthcheckHandler(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		resp := test.GET(ctx, &test.Request{URL: "/healthcheck"})

		is.Equal(http.StatusOK, resp.Code)
	})
}
