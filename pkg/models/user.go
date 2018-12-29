package models

import (
	"time"

	"github.com/google/uuid"
)

// User is the user.
type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"modified_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`

	Email            string `json:"email"`
	EmailPublic      bool   `json:"email_public"`
	IsEmailConfirmed bool   `json:"is_email_confirmed"`

	FullName    string `json:"full_name"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Website     string `json:"website"`

	LanguageUsedPublic bool      `json:"languages_used_public"`
	LastHeartBeat      time.Time `json:"last_heartbeat"`

	LastPlugin     string `json:"last_plugin"`
	LastPluginName string `json:"last_plugin_name"`

	LastProject string `json:"last_project"`

	Location string `json:"location"`
	TimeZone string `json:"timezone"`

	LoggedTimePublic bool `json:"logged_time_public"`

	Photo       string `json:"photo"`
	PhotoPublic bool   `json:"photo_public"`

	// Useless Fields
	HasPremiumFeatures bool   `json:"has_premium_features"`
	IsHireable         bool   `json:"is_hireable"`
	Plan               string `json:"plan"`
}
