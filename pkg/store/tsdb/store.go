package tsdb

import (
	"time"

	"github.com/dyweb/sundial/pkg/models"
)

// Store is the interface of time series data store.
type Store interface {
	// WriteHeartBeats write the heartbeats to db.
	WriteHeartBeats(username string, heartBeats []models.HeartBeatFrontModel) error
	// QueryHeartBeats gets all heartbeats in a given time range.
	QueryHeartBeats(begin time.Time, end time.Time) ([]models.HeartBeatFrontModel, error)
}
