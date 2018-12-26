package datastore

import (
	"github.com/google/uuid"

	"github.com/dyweb/sundial/pkg/models"
)

const projectTable = "projects"

func (ds *datastore) GetProject(UUID uuid.UUID) (*models.Project, error) {
	var project = &models.Project{}
	err := ds.First(project, "id = ?", UUID).Error
	return project, err
}

func (ds *datastore) CreateProject(project *models.Project) error {
	return ds.Create(project).Error
}
