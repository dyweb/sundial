package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	UUID            uuid.UUID      `json:"id" gorm:"PRIMARY_KEY"`
	Name            string         `json:"name"`
	HTMLEscapedName string         `json:"html_escaped_name"`
	Privacy         string         `json:"privacy"`
	Repository      NotImplemented `json:"repository"`
	URL             string         `json:"url"`
}

type NotImplemented struct {
}
