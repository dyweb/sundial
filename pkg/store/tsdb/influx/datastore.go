package influx

import (
	"log"

	client "github.com/influxdata/influxdb/client/v2"
)

type DataStore struct {
	client.Client
}

func New(source, username, pwd string) *DataStore {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     source,
		Username: username,
		Password: pwd,
	})
	if err != nil {
		log.Fatal("Failed to create the tsdb datastore: %v", err)
	}
	return &DataStore{
		Client: c,
	}
}
