package manager_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/manager"
	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/store"
	"github.com/thoas/observr/store/models"
	"github.com/thoas/observr/test"
)

func TestManager_User_UserCreate(t *testing.T) {
	test.Setup(func(ctx context.Context) {
		is := require.New(t)

		u := &models.User{
			Username: "thoas",
			Email:    "florent@ulule.com",
			Password: "$ecret",
		}

		err := store.CreateUser(ctx, u)
		is.Nil(err)

		user, err := manager.CreateUser(ctx, &payloads.UserCreate{
			Username: "thoas",
			Email:    "florent@ulule.com",
			Password: "$ecret",
		})

		is.NotNil(err)
		is.Nil(user)

		httpError := err.(failure.HTTPError)

		is.Equal(1, httpError.Errors.Len())
		is.Equal(httpError.Errors[0].Kind(), "AlreadyExistsError")
		is.Equal(httpError.Errors[0].Fields(), []string{"email"})

		user, err = manager.CreateUser(ctx, &payloads.UserCreate{
			Username: "thoas",
			Email:    "florent+2@ulule.com",
			Password: "$ecret",
		})

		is.NotNil(err)
		is.Nil(user)

		httpError = err.(failure.HTTPError)

		is.Equal(1, httpError.Errors.Len())
		is.Equal(httpError.Errors[0].Kind(), "AlreadyExistsError")
		is.Equal(httpError.Errors[0].Fields(), []string{"username"})
	})
}
