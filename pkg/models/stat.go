package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Stat is a stat of a given time period.
type Stat struct {
	gorm.Model
	UUID             uuid.UUID  `json:"id"`
	TotalSeconds     int64      `json:"total_seconds"`
	Projects         []Workunit `json:"projects"`
	Languages        []Workunit `json:"languages"`
	Editors          []Workunit `json:"editors"`
	OperatingSystems []Workunit `json:"operating_systems"`
	Dependencies     []Workunit `json:"dependencies"`
	BestDay          struct {
		Date         time.Time `json:"date"`
		TotalSeconds int64     `json:"total_seconds"`
	} `json:"best_day"`
	Range StatRange `json:"range"`
	//Holidays is number of days in this range with no coding time logged
	Holidays              int64 `json:"holidays"`
	DaysIncludingHolidays int64 `json:"days_including_holidays"`
	DaysMinusHolidays     int64 `json:"days_minus_holidays"`
	//Status is "ok"
	Status string `json:"status"`
	//is_already_updating is true if these stats are being updated in the background
	//TODO: currently always false
	IsAlreadyUpdating bool `json:"is_already_updating"`
	//is_coding_activity_visible is true if this user's coding activity is publicly visible
	//TODO: currently always true
	IsCodingActivityVisible bool `json:"is_coding_activity_visible"`
	//is_other_usage_visible is true if this user's languages, editors, and operating system stats are publicly visible
	//TODO: currently always true
	IsOtherUsageVisible bool `json:"is_other_usage_visible"`
	//is_stuck is true if these stats got stuck while processing and will be recalculated in the background
	//TODO: currently always false
	IsStuck bool `json:"is_stuck"`
	//is_up_to_date is true if these stats are up to date
	//TODO: currently always true
	IsUpToDate bool `json:"is_up_to_date"`
	//start is start of this time range as ISO 8601 UTC datetime
	Start int64 `json:"start"`
	//end is end of this time range as ISO 8601 UTC datetime
	End int64 `json:"end"`
	//timezone is timezone used in Olson Country/Region format
	Timezone string `json:"timezone"`
	//timeout is value of the user's timeout setting in seconds
	Timeout int64 `json:"timeout"`
	//writes_only is status of the user's writes_only setting
	WritesOnly bool `json:"writes_only"`
	//user_id is unique id of this user
	UserID uuid.UUID `json:"user_id"`
	//username is public username for this user
	Username string `json:"username"`
	//created_at is time when these stats were created in ISO 8601 format
	CreatedAt time.Time `json:"created_at"`
	//modified_at is time when these stats were last updated in ISO 8601 format
	//TODO: currently always equal to CreatedAt
	ModifiedAt time.Time `json:"modified_at"`
}

// Workunit is a unit of work,
// used in stats to present time spent on a piece of work.
type Workunit struct {
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	TotalSeconds int64   `json:"total_seconds"`
}

//StatRange is the range of time a `Stat` is about.
type StatRange = string

const (
	//StatRangeLast7Days is about last 7 days
	StatRangeLast7Days = "last_7_days"
	// StatRangeLast30Days is about last 30 days
	StatRangeLast30Days = "last_30_days"
	//StatRangeLast6Months is about last 6 months
	StatRangeLast6Months = "last_6_months"
	//StatRangeLastYear is about last year
	StatRangeLastYear = "last_year"
)
