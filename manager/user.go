package manager

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/thoas/observr/failure"
	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/store"
	"github.com/thoas/observr/store/models"
)

func CreateUser(ctx context.Context, payload *payloads.UserCreate) (*models.User, error) {
	email := strings.ToLower(payload.Email)

	exists, err := store.UserEmailExists(ctx, email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, failure.AlreadyExists([]string{"email"})
	}

	username := payload.Username

	exists, err = store.UsernameExists(ctx, username)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, failure.AlreadyExists([]string{"username"})
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(pwd),
		Email:    email,
	}

	err = store.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
