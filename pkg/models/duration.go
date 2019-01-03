package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Duration models a piece of duration
type Duration struct {
	Project  string    `json:"project"`
	Time     time.Time `json:"time"`     //begin of the duration
	Duration int       `json:"duration"` //length of the duration
}

//StoredDuration is saved in gorm
type StoredDuration struct {
	gorm.Model
	Duration
	Branch string
}

//DurationResponse models response message of GET /users/:user/durations
type DurationResponse struct {
	Data     []Duration `json:"data"`
	Branches []string   `json:"branches"`
	Start    time.Time  `json:"start"`
	End      time.Time  `json:"end"`
	Timezone string     `json:"timezone"` //timezone used for this request in Olson Country/Region format
}
