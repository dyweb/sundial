package users

import (
	"context"

	"github.com/caicloud/nirvana/log"

	"github.com/dyweb/sundial/pkg/models"
)

// GetUser returns a user by username.
func GetUser(ctx context.Context, username string) (*models.User, error) {
	return &models.User{
		DisplayName: username,
	}, nil
}

// GetCurrentUser returns the current user.
func GetCurrentUser(ctx context.Context) (*models.User, error) {
	log.Info("Hello")
	return &models.User{
		DisplayName: "Current Name",
	}, nil
}
