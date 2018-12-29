package projects

import (
	"context"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/rdb"
)

// GetProjects returns the list of projects for the given user.
func GetProjects(ctx context.Context, username string) ([]models.Project, error) {
	ds := rdb.FromContext(ctx)
	return ds.GetProjects()
}

// GetCurrentProjects returns the list of projects for the current user.
func GetCurrentProjects(ctx context.Context) ([]models.Project, error) {
	return GetProjects(ctx, "")
}
