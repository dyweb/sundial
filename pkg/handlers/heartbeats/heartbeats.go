package heartbeats

import (
	"context"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/tsdb"
)

// POSTHeartBeat handles the post method for heartbeat.
func POSTHeartBeat(ctx context.Context, username string,
	heartbeats []models.HeartBeatFrontModel) ([]models.HeartBeatFrontModel, error) {
	tsstore := tsdb.FromContext(ctx)
	err := tsstore.WriteHeartBeats(username, heartbeats)
	return heartbeats, err
}

// POSTCurrentHeartBeat handles the post method for heartbeat.
func POSTCurrentHeartBeat(ctx context.Context, heartbeats []models.HeartBeatFrontModel) ([]models.HeartBeatFrontModel, error) {
	return POSTHeartBeat(ctx, "current", heartbeats)
}
