package store

import (
	"github.com/dyweb/sundial/pkg/models"
)

// Store is the interface of data store.
type Store interface {
	// GetProject creates a new project.
	GetProject(id int64) (*models.Project, error)
	// CreateProject creates a new project.
	CreateProject(*models.Project) error
}
