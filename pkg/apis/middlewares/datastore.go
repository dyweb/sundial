package middlewares

import (
	"context"

	def "github.com/caicloud/nirvana/definition"
	"github.com/dyweb/sundial/pkg/store/rdb"
)

// MakeStore is a middleware function that initializes the Datastore and attaches to
// the context of every http.Request.
func MakeStore(s rdb.Store) def.Middleware {
	return func(ctx context.Context, next def.Chain) error {
		return next.Continue(rdb.ToContext(ctx, s))
	}
}
