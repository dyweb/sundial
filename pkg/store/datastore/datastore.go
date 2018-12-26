package datastore

import (
	"time"

	"github.com/caicloud/nirvana/log"
	"github.com/jinzhu/gorm"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store"
)

// datastore is an implementation of a model.Store built on top
// of the sql/database driver with a relational database backend.
type datastore struct {
	*gorm.DB

	driver string
	config string
}

// New creates a database connection for the given driver and datasource
// and returns a new Store.
func New(driver, config string) store.Store {
	ds := &datastore{
		DB:     open(driver, config),
		driver: driver,
		config: config,
	}
	ds.AutoMigrate(&models.Project{})
	return ds
}

// open opens a new database connection with the specified
// driver and connection string and returns a store.
func open(driver, config string) *gorm.DB {
	db, err := gorm.Open(driver, config)
	if err != nil {
		log.Errorln(err)
		log.Fatalln("database connection failed")
	}

	if err := pingDatabase(db); err != nil {
		log.Errorln(err)
		log.Fatalln("database ping attempts failed")
	}
	return db
}

// helper function to ping the database with backoff to ensure
// a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *gorm.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.DB().Ping()
		if err == nil {
			return
		}
		log.Infof("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return
}
