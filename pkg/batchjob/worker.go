package batchjob

import (
	"github.com/robfig/cron"

	"github.com/dyweb/sundial/pkg/store/rdb"
	"github.com/dyweb/sundial/pkg/store/tsdb"
)

type Manager struct {
	RStore  rdb.Store
	TSStore tsdb.Store
	Cron    cron.Cron
}

func NewManager(rdbStore rdb.Store, tsStore tsdb.Store) *Manager {
	return &Manager{
		RStore:  rdbStore,
		TSStore: tsStore,
		Cron:    cron.New(),
	}
}

func (m Manager) SummarizeDuration() {

}
