package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Project is the project that could be contributed by a user.
type Project struct {
	gorm.Model
	UUID            uuid.UUID      `json:"id"`
	Name            string         `json:"name"`
	HTMLEscapedName string         `json:"html_escaped_name"`
	Privacy         string         `json:"privacy"`
	Repository      NotImplemented `json:"repository"`
	URL             string         `json:"url"`
}

type NotImplemented struct {
}
