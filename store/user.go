package store

import (
	"context"

	"github.com/thoas/observr/store/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	user.ID = UUID()

	return Sync(ctx, usersCreateQuery, user)
}
