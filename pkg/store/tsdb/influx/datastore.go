package influx

import (
	"fmt"
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

// QueryHeartBeats gets all heartbeats in a given time range.
func (ds DataStore) QueryHeartBeats(begin time.Time, end time.Time) ([]models.HeartBeatFrontModel, error) {
	queryString := fmt.Sprintf("SELECT * FROM %s WHERE time < %s AND time > %s", constants.MeasurementName, begin.Format(time.RFC3339), end.Format(time.RFC3339))
	responses, err := ds.Query(client.NewQuery(queryString, ds.DBName, "s"))
	if err != nil {
		return nil, err
	}
	heartbeats := []models.HeartBeatFrontModel{}
	for _, result := range responses.Results {
		for _, row := range result.Series {
			for _, onepoint := range row.Values {
				hb := models.HeartBeatFrontModel{}
				//TODO: currently some fields of HB is not stored into influxDB
				//so when we retrieve it back we could only get empty placeholders.
				//someday we may be able to store them.
				hb.Dependencies = []string{}

				for i := 0; i < len(onepoint); i++ {
					switch row.Columns[i] {
					case "time":
						hb.Time = onepoint[i].(float64)
					case "user":
						//TODO: heartbeat model does not save username now
					case "entity":
						hb.Entity = onepoint[i].(string)
					case "type":
						hb.Type = onepoint[i].(models.HeartBeatType)
					case "category":
						hb.Category = onepoint[i].(models.HeartBeatCategory)
					case "project":
						hb.Project = onepoint[i].(string)
					case "branch":
						hb.Branch = onepoint[i].(string)
					case "language":
						hb.Language = onepoint[i].(string)
					case "is_write":
						hb.IsWrite = onepoint[i].(bool)
					case "lines":
						hb.Lines = onepoint[i].(int64)
					case "lineno":
						hb.Lineno = onepoint[i].(int64)
					case "cursorpos":
						hb.Cursorpos = onepoint[i].(int64)
					default:
						//do nothing
					}
				}
				heartbeats = append(heartbeats, hb)
			}
		}
	}
	return heartbeats, nil
}
