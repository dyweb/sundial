package models

import "time"

type hb struct {
	// Entity defines entity heartbeat is logging time against, such as an absolute file path or domain.
	Entity string `json:"entity"`
	// Type is the type of entity.
	Type HeartBeatType `json:"type"`
	// Category is the category for this activity.
	Category HeartBeatCategory `json:"category"`
	// Project is the project name (optional).
	Project string `json:"project,omitempty"`
	// Branch is the branch name (optional).
	Branch string `json:"branch,omitempty"`
	// Language is the language name (optional).
	Language string `json:"language,omitempty"`
	// Dependencies is comma separated list of dependencies detected from entity file (optional).
	Dependencies []string `json:"dependencies,omitempty"`
	// Lines is the total number of lines in the entity (when entity type is file).
	Lines int64 `json:"lines,omitempty"`
	// Lineno is the current line row number of cursor (optional).
	Lineno int64 `json:"lineno,omitempty"`
	// Cursorpos is the current cursor column position (optional).
	Cursorpos int64 `json:"cursorpos,omitempty"`
	// IsWrite defines whether this heartbeat was triggered from writing to a file.
	IsWrite bool `json:"is_write,omitempty"`
}

// HeartBeatFrontModel describes a HeartBeat entry for frontend.
type HeartBeatFrontModel struct {
	hb
	// Time is UNIX epoch timestamp; numbers after decimal point are fractions of a second.
	Time float64 `json:"time"`
}

// HeartBeat describes a HeartBeat entry.
type HeartBeat struct {
	hb
	// Time is time that the heartbeat is created.
	Time time.Time `json:"time"`
}

// HeartBeatType is the type of entity.
type HeartBeatType = string

const (
	// HeartBeatTypeFile is the file type of entity.
	HeartBeatTypeFile HeartBeatType = "file"
	// HeartBeatTypeApp is the app type of entity.
	HeartBeatTypeApp HeartBeatType = "app"
	// HeartBeatTypeDomain is the domain type of entity.
	HeartBeatTypeDomain HeartBeatType = "domain"
)

// HeartBeatCategory is the category for this activity.
type HeartBeatCategory = string

const (
	// HeartBeatCategoryCoding is the coding category for this activity.
	HeartBeatCategoryCoding HeartBeatCategory = "coding"
	// HeartBeatCategoryBuilding is the building category for this activity.
	HeartBeatCategoryBuilding HeartBeatCategory = "building"
	// HeartBeatCategoryIndexing is the indexing category for this activity.
	HeartBeatCategoryIndexing HeartBeatCategory = "indexing"
	// HeartBeatCategoryDebugging is the debugging category for this activity.
	HeartBeatCategoryDebugging HeartBeatCategory = "debugging"
	// HeartBeatCategoryBrowsing is the browsing category for this activity.
	HeartBeatCategoryBrowsing HeartBeatCategory = "browsing"
	// HeartBeatCategoryRunningTests is the running tests category for this activity.
	HeartBeatCategoryRunningTests HeartBeatCategory = "running tests"
	// HeartBeatCategoryWrittingTests is the writing tests category for this activity.
	HeartBeatCategoryWrittingTests HeartBeatCategory = "writing tests"
	// HeartBeatCategoryManualTesting is the manual testing category for this activity.
	HeartBeatCategoryManualTesting HeartBeatCategory = "manual testing"
	// HeartBeatCategoryCodeReviewing is the code reviewing category for this activity.
	HeartBeatCategoryCodeReviewing HeartBeatCategory = "code reviewing"
	// HeartBeatCategoryDesigning is the designing category for this activity.
	HeartBeatCategoryDesigning HeartBeatCategory = "designing"
)
