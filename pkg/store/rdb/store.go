package rdb

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dyweb/sundial/pkg/models"
)

// Store is the interface of data store.
type Store interface {
	// GetProject creates a new project.
	GetProject(UUID uuid.UUID) (*models.Project, error)
	// CreateProject creates a new project.
	CreateProject(*models.Project) error
	// GetProjects gets all projects.
	GetProjects() ([]models.Project, error)
}
