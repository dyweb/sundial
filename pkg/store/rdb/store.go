package rdb

import (
	"time"

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
	// GetDuration gets all durations in that day, that project and that branches.
	// user: required.
	// day: required. Durations will be returned from 12am until 11:59pm in user's timezone for this day.
	// project: optional. nil means all projects.
	// branches: optional. nil means all branches. empty array means no branches (and empty result).
	GetDurations(user string, date time.Time, project string, branches []string) ([]models.StoredDuration, error)
}
