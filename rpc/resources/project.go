package resources

import (
	"context"
	"time"

	"github.com/ulule/deepcopier"

	"github.com/thoas/observr/store/models"
)

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProject(ctx context.Context, project *models.Project) (*Project, error) {
	var resource Project

	err := deepcopier.Copy(project).To(&resource)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
