package influx

import (
	"log"
	"time"

	"github.com/dyweb/sundial/pkg/constants"
	"github.com/dyweb/sundial/pkg/models"
	client "github.com/influxdata/influxdb/client/v2"
)

type DataStore struct {
	client.Client
	DBName string
}

func New(source, username, pwd, dbName string) *DataStore {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     source,
		Username: username,
		Password: pwd,
	})
	if err != nil {
		log.Fatalf("Failed to create the tsdb datastore: %v", err)
	}
	return &DataStore{
		Client: c,
		DBName: dbName,
	}
}

// WriteHeartBeats write the heartbeats to db.
func (ds DataStore) WriteHeartBeats(username string, heartBeats []models.HeartBeatFrontModel) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  ds.DBName,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	for _, hb := range heartBeats {
		tags := map[string]string{
			"user":     username,
			"entity":   hb.Entity,
			"type":     hb.Type,
			"category": hb.Category,
			"project":  hb.Project,
			"branch":   hb.Branch,
			"language": hb.Language,
			// dependencies is not supported because influxdb does not support slice.
		}
		fields := map[string]interface{}{
			"is_write":  hb.IsWrite,
			"lines":     hb.Lines,
			"lineno":    hb.Lineno,
			"cursorpos": hb.Cursorpos,
		}

		pt, err := client.NewPoint(
			constants.MeasurementName,
			tags,
			fields,
			toTime(hb.Time),
		)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
	}

	return ds.Write(bp)
}

func toTime(t float64) time.Time {
	return time.Unix(int64(t), 0)
}
