package store

import (
	"github.com/dyweb/sundial/pkg/models"
	"github.com/google/uuid"
)

// Store is the interface of data store.
type Store interface {
	// GetProject creates a new project.
	GetProject(UUID uuid.UUID) (*models.Project, error)
	// CreateProject creates a new project.
	CreateProject(*models.Project) error
}
