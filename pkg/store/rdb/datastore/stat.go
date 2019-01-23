package datastore

import (
	"github.com/dyweb/sundial/pkg/models"
)

//CreateStat saves a stat.
func (ds *datastore) CreateStat(stat *models.Stat) error {
	return ds.Create(stat).Error
}
