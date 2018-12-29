package models

import (
	"github.com/jinzhu/gorm"
)

// Message describes a message entry.
type Message struct {
	gorm.Model
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
