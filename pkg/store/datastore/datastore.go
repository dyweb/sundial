package datastore

import (
	"database/sql"
	"time"

	"github.com/caicloud/nirvana/log"
	"github.com/russross/meddler"

	"github.com/dyweb/sundial/pkg/store"
)

// datastore is an implementation of a model.Store built on top
// of the sql/database driver with a relational database backend.
type datastore struct {
	*sql.DB

	driver string
	config string
}

// New creates a database connection for the given driver and datasource
// and returns a new Store.
func New(driver, config string) store.Store {
	return &datastore{
		DB:     open(driver, config),
		driver: driver,
		config: config,
	}
}

// From returns a Store using an existing database connection.
func From(db *sql.DB) store.Store {
	return &datastore{DB: db}
}

// open opens a new database connection with the specified
// driver and connection string and returns a store.
func open(driver, config string) *sql.DB {
	db, err := sql.Open(driver, config)
	if err != nil {
		log.Errorln(err)
		log.Fatalln("database connection failed")
	}
	if driver == "mysql" {
		// per issue https://github.com/go-sql-driver/mysql/issues/257
		db.SetMaxIdleConns(0)
	}

	setupMeddler(driver)

	if err := pingDatabase(db); err != nil {
		log.Errorln(err)
		log.Fatalln("database ping attempts failed")
	}
	return db
}

// helper function to ping the database with backoff to ensure
// a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		log.Infof("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return
}

// helper function to setup the meddler default driver
// based on the selected driver name.
func setupMeddler(driver string) {
	switch driver {
	case "sqlite3":
		meddler.Default = meddler.SQLite
	case "mysql":
		meddler.Default = meddler.MySQL
	case "postgres":
		meddler.Default = meddler.PostgreSQL
	}
}
