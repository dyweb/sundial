package datastore

import (
	"time"

	"github.com/dyweb/sundial/pkg/models"
)

// GetDuration gets all durations in that day, that project and that branches.
// user: required.
// day: required. Durations will be returned from 12am until 11:59pm in user's timezone for this day.
// project: optional. nil means all projects.
// branches: optional. nil means all branches. empty array means no branches (and empty result).
func (ds *datastore) GetDurations(user string, date time.Time, project string, branches []string) ([]models.StoredDuration, error) {
	result := []models.StoredDuration{}
	if err := ds.
		Where("Time between ? and ?", date, date.AddDate(0, 0, 1)).
		Where("Project = ?", project).
		Where("Branch in (?)", branches).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
