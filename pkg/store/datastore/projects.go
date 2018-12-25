package datastore

import (
	"github.com/russross/meddler"

	"github.com/dyweb/sundial/pkg/models"
)

const projectTable = "projects"

func (db *datastore) GetProject(id int64) (*models.Project, error) {
	var project = new(models.Project)
	var err = meddler.Load(db, projectTable, project, id)
	return project, err
}

func (db *datastore) CreateProject(project *models.Project) error {
	return meddler.Insert(db, projectTable, project)
}
