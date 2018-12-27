package middlewares

import (
	def "github.com/caicloud/nirvana/definition"

	"github.com/dyweb/sundial/pkg/store/rdb"
)

// Middlewares returns a list of middlewares.
func Middlewares() []def.Middleware {
	return []def.Middleware{}
}

// WithStore returns a list with store middleware.
func WithStore(s rdb.Store) []def.Middleware {
	return []def.Middleware{MakeStore(s)}
}
