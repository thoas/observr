package store

import (
	"context"
	"fmt"

	"github.com/dchest/uniuri"
	"github.com/thoas/observr/store/models"
)

func CreateProject(ctx context.Context, project *models.Project) error {
	project.ID = UUID()
	project.ApiKey = fmt.Sprintf("obs_%s", uniuri.New())

	return Sync(ctx, projectsCreateQuery, project)
}
