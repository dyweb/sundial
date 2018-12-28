package tsdb

import "github.com/dyweb/sundial/pkg/models"

// Store is the interface of time series data store.
type Store interface {
	// WriteHeartBeats write the heartbeats to db.
	WriteHeartBeats(username string, heartBeats []models.HeartBeatFrontModel) error
}
