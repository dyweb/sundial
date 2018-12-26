package store

import (
	"golang.org/x/net/context"
)

type ContextKey string

const (
	// Key is the key of store in the context.
	Key ContextKey = "store"
)

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) Store {
	return c.Value(Key).(Store)
}

// ToContext adds the Store to this context if it supports
// the Setter interface.
func ToContext(c context.Context, store Store) context.Context {
	return context.WithValue(c, Key, store)
}
