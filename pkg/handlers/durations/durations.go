package durations

import (
	"context"
	"strings"
	"time"

	"github.com/caicloud/nirvana/log"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/rdb"
)

// GetDurations returns a user's durations
func GetDurations(ctx context.Context, user string, dateString string, project string, branches string) (*models.DurationResponse, error) {
	log.Info("GetDurations", user, dateString, project, branches)
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, err
	}

	ds := rdb.FromContext(ctx)
	response := models.DurationResponse{}
	resultBranches := map[string]bool{} //used as sets

	durations, err := ds.GetDurations(user, date, project, strings.Split(branches, ","))
	if err != nil {
		return nil, err
	}

	for _, duration := range durations {
		response.Data = append(response.Data, duration.Duration)
		if _, ok := resultBranches[duration.Branch]; ok {
			//do nothing
		} else {
			response.Branches = append(response.Branches, duration.Branch)
			resultBranches[duration.Branch] = true
		}
	}

	response.Start = date
	response.End = date.AddDate(0, 0, 1)
	//TODO: read from User
	response.Timezone = "America/Los_Angeles"

	return &response, nil
}

// GetCurrentUserDurations returns the current user's durations.
func GetCurrentUserDurations(ctx context.Context, dateString string, project string, branches string) (*models.DurationResponse, error) {
	log.Info("Hello")
	return GetDurations(ctx, "current", dateString, project, branches)
}
