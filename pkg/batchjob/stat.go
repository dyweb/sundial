package batchjob

import (
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/rdb"
	"github.com/dyweb/sundial/pkg/store/tsdb"
)

//StatWorker collects heartbeats from tsdb and writes stat to rdb
//collects heartbeats from `startTime` with a given range length
func StatWorker(rdbStore rdb.Store, tsStore tsdb.Store, beginTime time.Time, statRange models.StatRange) error {
	hbs, err := tsStore.QueryHeartBeats(beginTime, beginTime.Add(toDuration(statRange)))
	if err != nil {
		return err
	}
	stat, err := calculateStat(hbs, beginTime, statRange)
	if err != nil {
		return err
	}
	err = rdbStore.CreateStat(stat)
	if err != nil {
		return err
	}
	return nil
}

//aggregation rules:
//each heartbeat contains 1 jiffy, whose size is now hardcoded into 1 minute.
//if two heartbeats are closer than 1 jiffy, merge them.
//a heartbeat contributes to its project, language, editor, os and dependencies.
//assume `hbs` is in ascending time order.
func calculateStat(hbs []models.HeartBeatFrontModel, beginTime time.Time, statRange models.StatRange) (*models.Stat, error) {
	//TODO: make jiffy configurable, e.g. argument of the cmd
	jiffy := 60.0 //seconds for one minute

	stat := models.Stat{}
	// part 1: hardcoded fields
	stat.IsAlreadyUpdating = false
	stat.IsCodingActivityVisible = true
	stat.IsOtherUsageVisible = true
	stat.IsStuck = false
	stat.IsUpToDate = true
	stat.Status = "ok"
	// these 3 fields are set empty for now, see [process.md](docs/process.md)
	stat.Editors = []models.Workunit{}
	stat.OperatingSystems = []models.Workunit{}
	stat.Dependencies = []models.Workunit{}
	// part 2: fields that should be read from User
	//TODO: fix them when we introduce a user
	stat.Timeout = 60
	stat.Timezone = "America/Los_Angeles"
	stat.WritesOnly = false
	stat.UserID = uuid.UUID{} //now just a random uuid
	stat.Username = "currentUser"

	// part 3: aggregations
	stat.DaysIncludingHolidays = int64(toDuration(statRange)) / int64(time.Hour*24)

	if len(hbs) > 0 {
		totalSeconds := float64(0)
		prevEnd := float64(0)
		// days: we define day as 24hour-aligned times.
		// Not that fancy considering leap seconds or so.
		workdays, prevDay := int64(0), dayFromFloatUnix(hbs[0].Time)
		bestDay, bestDaySeconds, thisDaySeconds := prevDay, float64(0), float64(0)

		projects := Workunits{units: make(map[string]WorkunitWithEndtime)}
		languages := Workunits{units: make(map[string]WorkunitWithEndtime)}
		for _, hb := range hbs {
			upsertWorkunit(&projects, hb.Project, hb.Time, jiffy)
			upsertWorkunit(&languages, hb.Language, hb.Time, jiffy)
			elapsed, newEnd := mergeIntervals(prevEnd, hb.Time, jiffy)
			totalSeconds += elapsed
			prevEnd = newEnd

			day := dayFromFloatUnix(hb.Time)
			if day != prevDay {
				workdays++

				if thisDaySeconds > bestDaySeconds {
					bestDay, bestDaySeconds = day, thisDaySeconds
				}

				prevDay = day
				thisDaySeconds = 0.0

			} else {
				thisDaySeconds += elapsed
			}
		}
		stat.Projects = percentedWorkunits(projects)
		stat.Languages = percentedWorkunits(languages)

		stat.DaysMinusHolidays = workdays + 1
		stat.Holidays = stat.DaysIncludingHolidays - stat.DaysMinusHolidays
		stat.BestDay.Date = bestDay
		stat.BestDay.TotalSeconds = int64(bestDaySeconds)
		stat.Start = int64(hbs[0].Time)
		stat.End = int64(hbs[len(hbs)-1].Time + jiffy)
		stat.TotalSeconds = int64(totalSeconds)
	} else {
		//if there are no hbs in this range
		stat.Projects = []models.Workunit{}
		stat.Languages = []models.Workunit{}
		stat.DaysMinusHolidays = 0
		stat.Holidays = stat.DaysIncludingHolidays
		stat.BestDay.Date = beginTime
		stat.BestDay.TotalSeconds = 0
		stat.Start = beginTime.Unix()
		stat.End = beginTime.Add(toDuration(statRange)).Unix()
		stat.TotalSeconds = 0
	}

	// part 4: metadata
	stat.Range = statRange
	stat.CreatedAt = time.Now()
	stat.ModifiedAt = stat.CreatedAt

	return &stat, nil
}

// 3 cases
// 1. start + jiffy < old endtime -> this hb contributes nothing
// 2. start < old endtime < time + jiffy -> this hb contributes less than a jiffy
// 3. old endtime < start -> there is a discontinuity of work, and hb contributes a jiffy
func mergeIntervals(prevEnd float64, start float64, jiffy float64) (elapsed float64, newEnd float64) {
	if start+jiffy < prevEnd {
		return 0, prevEnd
	} else if start < prevEnd {
		return start + jiffy - prevEnd, start + jiffy
	} else {
		return jiffy, start + jiffy
	}
}

func upsertWorkunit(workunits *Workunits, name string, start float64, jiffy float64) {
	/* if such name does not exist, create one ending in 1970 */
	if _, ok := workunits.units[name]; !ok {
		workunits.units[name] = WorkunitWithEndtime{
			Workunit: models.Workunit{
				Name:         name,
				Percent:      0,
				TotalSeconds: 0,
			},
			Endtime: 0,
		}
	}

	/* update it */
	wu, _ := workunits.units[name]
	elapsed, newEnd := mergeIntervals(wu.Endtime, start, jiffy)
	wu.Endtime = newEnd
	wu.TotalSeconds += int64(elapsed)
	workunits.totalSeconds += elapsed
}

//WorkunitWithEndtime is a temporal struct for aggregation
type WorkunitWithEndtime struct {
	models.Workunit
	Endtime float64
}

//calculate percents for each wu, and return apporpriate type for output.
func percentedWorkunits(wus Workunits) []models.Workunit {
	outputs := make([]models.Workunit, len(wus.units))
	for _, wu := range wus.units {
		wu.Percent = float64(wu.TotalSeconds) / wus.totalSeconds
		outputs = append(outputs, wu.Workunit)
	}
	return outputs
}

//Workunits is a temporal struct for aggregation
type Workunits struct {
	units        map[string]WorkunitWithEndtime
	totalSeconds float64
}

//toDuration converts semantics of statRange into time.Duration.
//FIXME: it is an approximation.
//To get real duration, we have to decide whether use calendar date or duration date
//and have to have another argument of beginTime.
func toDuration(statRange models.StatRange) time.Duration {
	switch statRange {
	case models.StatRangeLast7Days:
		return time.Hour * 24 * 7
	case models.StatRangeLast30Days:
		return time.Hour * 24 * 30
	case models.StatRangeLast6Months:
		return time.Hour * 24 * 30 * 6
	case models.StatRangeLastYear:
		return time.Hour * 24 * 365
	}
	// will not happen
	log.Fatalln("unexpected case of statRange")
	return 0
}

func dayFromFloatUnix(unix float64) time.Time {
	secs := int64(unix)
	nsecs := int64((unix - float64(secs)) * 1000000000)
	return time.Unix(secs, nsecs).Truncate(time.Hour * 24)
}
