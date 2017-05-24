package manager

import (
	"context"

	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/store/models"
)

func CreateUser(ctx context.Context, payload *payloads.UserCreatePayload) (*models.User, error) {
	return nil, nil
}
