package apis

import "github.com/dyweb/sundial/pkg/store/rdb/datastore"

var (
	// TODO(gaocegege): This should be configurable.
	ds = datastore.New("sqlite3", "sundial.sqlite")
)
