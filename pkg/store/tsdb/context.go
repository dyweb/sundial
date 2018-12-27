package tsdb

import (
	"golang.org/x/net/context"
)

type ContextKey string

const (
	// Key is the key of store in the context.
	key ContextKey = "tsdbstore"
)

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

// ToContext adds the Store to this context if it supports
// the Setter interface.
func ToContext(c context.Context, store Store) context.Context {
	return context.WithValue(c, key, store)
}
