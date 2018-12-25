package models

import (
	"github.com/google/uuid"
)

type Project struct {
	ID              int64          `meddler:"project_id,pk"`
	UUID            uuid.UUID      `json:"id" meddler:"project_uuid,pk"`
	HTMLEscapedName string         `json:"html_escaped_name" meddler:"project_html_escaped_name"`
	Privacy         string         `json:"privacy" meddler:"project_privacy"`
	Repository      NotImplemented `json:"repository" meddler:"project_repository"`
	URL             string         `json:"url" meddler:"project_url"`
}

type NotImplemented struct {
}
