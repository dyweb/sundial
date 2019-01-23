package batchjob

import (
	"errors"
	"log"
	"time"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/rdb"
	"github.com/dyweb/sundial/pkg/store/tsdb"
)

//StatWorker collects heartbeats from tsdb and writes stat to rdb
//collects heartbeats from `startTime` with a given range length
func StatWorker(rdbStore rdb.Store, tsStore tsdb.Store, beginTime time.Time, statRange models.StatRange) {
	hbs, err := tsStore.QueryHeartBeats(beginTime, beginTime.Add(toDuration(statRange)))
	if err != nil {
		log.Fatalln(err)
	}
	stat, err := calculateStat(hbs, beginTime, statRange)
	if err != nil {
		log.Fatalln(err)
	}
	err = rdbStore.CreateStat(stat)
	if err != nil {
		log.Fatalln(err)
	}
}

//aggregation rules:
//each heartbeat contains 1 jiffy, whose size is now hardcoded into 1 minute.
//if two heartbeats are closer than 1 jiffy, merge them.
//a heartbeat contributes to its project, language, editor, os and dependencies.
func calculateStat(hbs []models.HeartBeatFrontModel, beginTime time.Time, statRange models.StatRange) (*models.Stat, error) {
	jiffy := time.Minute
	//TODO
	return nil, errors.New("not implemented")
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
}
