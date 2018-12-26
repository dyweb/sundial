package apis

import "github.com/dyweb/sundial/pkg/store/datastore"

var (
	ds = datastore.New("sqlite3", "sundial.sqlite")
)
