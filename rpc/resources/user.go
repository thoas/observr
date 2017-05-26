package resources

import (
	"context"
	"time"

	"github.com/thoas/observr/store/models"
	"github.com/ulule/deepcopier"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(ctx context.Context, user *models.User) (*User, error) {
	var resource User

	err := deepcopier.Copy(user).To(&resource)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
