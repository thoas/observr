package handlers_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thoas/observr/test"
)

func TestHandlers_User_UserCreateHandler(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		is.Equal(1, 1)
	})
}
