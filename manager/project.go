package manager

import (
	"context"

	"github.com/thoas/observr/rpc/payloads"
	"github.com/thoas/observr/store"
	"github.com/thoas/observr/store/models"
)

func CreateProject(ctx context.Context, payload *payloads.ProjectCreate, user *models.User) (*models.Project, error) {
	project := &models.Project{
		Name:   payload.Name,
		URL:    payload.URL,
		UserID: user.ID,
	}

	err := store.CreateProject(ctx, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}
