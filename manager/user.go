package manager

import (
	"context"

	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/store"
	"github.com/thoas/observr/store/models"
)

func CreateUser(ctx context.Context, payload *payloads.UserCreate) (*models.User, error) {
	user := &models.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}

	err := store.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
