package users

import (
	"context"

	"github.com/caicloud/nirvana/log"
)

// User describes an user entry.
type User struct {
	DisplayName string `json:"display_name,omitempty"`
}

// GetUser returns a user by username.
func GetUser(ctx context.Context, username string) (*User, error) {
	return &User{
		DisplayName: username,
	}, nil
}

// GetCurrentUser returns the current user.
func GetCurrentUser(ctx context.Context) (*User, error) {
	log.Info("Hello")
	return &User{
		DisplayName: "Current Name",
	}, nil
}
