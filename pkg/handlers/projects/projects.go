package projects

import (
	"context"

	"github.com/dyweb/sundial/pkg/models"
)

// GetProjects returns the list of projects for the given user.
func GetProjects(ctx context.Context, username string) ([]models.Project, error) {
	// ds := rdb.FromContext(ctx)
	// err := ds.CreateProject(&models.Project{
	// 	Name: "test",
	// })
	return []models.Project{}, nil
}

// GetCurrentProjects returns the list of projects for the current user.
func GetCurrentProjects(ctx context.Context) ([]models.Project, error) {
	return []models.Project{}, nil
}
