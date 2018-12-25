package heartbeats

import (
	"context"

	"github.com/caicloud/nirvana/log"

	"github.com/dyweb/sundial/pkg/models"
)

// POSTHeartBeat handles the post method for heartbeat.
func POSTHeartBeat(ctx context.Context, username string, heartbeat models.HeartBeat) (models.HeartBeat, error) {
	log.Infof("HeartBeat: %v", heartbeat)
	return heartbeat, nil
}

// POSTCurrentHeartBeat handles the post method for heartbeat.
func POSTCurrentHeartBeat(ctx context.Context, heartbeats []models.HeartBeat) ([]models.HeartBeat, error) {
	log.Infof("%v", heartbeats)
	return heartbeats, nil
}
