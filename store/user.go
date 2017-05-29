package store

import (
	"context"

	"github.com/thoas/observr/store/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	user.ID = UUID()

	return Sync(ctx, usersCreateQuery, user)
}

func UserEmailExists(ctx context.Context, email string) (bool, error) {
	var result struct {
		Exists bool
		Email  string
	}

	result.Email = email

	err := Sync(ctx, usersEmailExistsQuery, &result)
	if err != nil {
		if IsErrNoRows(err) {
			return false, nil
		}

		return false, err
	}

	return result.Exists, nil
}

func UsernameExists(ctx context.Context, username string) (bool, error) {
	var result struct {
		Exists   bool
		Username string
	}

	result.Username = username

	err := Sync(ctx, usersUsernameExistsQuery, &result)
	if err != nil {
		if IsErrNoRows(err) {
			return false, nil
		}

		return false, err
	}

	return result.Exists, nil
}
