package handlers_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thoas/observr/test"
)

func TestHandlers_Base_HealthcheckHandler(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		resp := test.GET(ctx, "/healthcheck")

		is.Equal(200, resp.Code)
	})
}
