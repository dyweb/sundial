package main

import (
	"fmt"

	"github.com/robfig/cron"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/config"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/plugins/logger"
	"github.com/caicloud/nirvana/plugins/metrics"
	"github.com/caicloud/nirvana/plugins/reqlog"
	pversion "github.com/caicloud/nirvana/plugins/version"

	"github.com/dyweb/sundial/pkg/apis"
	"github.com/dyweb/sundial/pkg/apis/filters"
	"github.com/dyweb/sundial/pkg/apis/modifiers"
	"github.com/dyweb/sundial/pkg/batchjob"
	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/plugin"
	store "github.com/dyweb/sundial/pkg/store/plugin"
	rdbstore "github.com/dyweb/sundial/pkg/store/rdb/datastore"
	"github.com/dyweb/sundial/pkg/store/tsdb/influx"
	"github.com/dyweb/sundial/pkg/version"
)

func main() {

	//register cron for batchjob.
	//quick and dirty
	//TODO: shall have a registering mechaism, instead of hardcoding here.
	c := plugin.NewDefaultOption()
	s := rdbstore.New(c.RDBDriver, c.RDBSource)
	ts := influx.New(c.TSDBSource, c.TSDBUsername, c.TSDBPassword, c.TSDBName)
	crontab := cron.New()
	batchjob.AddCronJob(crontab, s, ts, models.StatRangeLast30Days)
	batchjob.AddCronJob(crontab, s, ts, models.StatRangeLast6Months)
	batchjob.AddCronJob(crontab, s, ts, models.StatRangeLast7Days)
	batchjob.AddCronJob(crontab, s, ts, models.StatRangeLastYear)

	// Print nirvana banner.
	fmt.Println(nirvana.Logo, nirvana.Banner)

	// Create nirvana command.
	cmd := config.NewNamedNirvanaCommand("sundial", config.NewDefaultOption())

	// Create plugin options.
	dbOption := store.NewDefaultOption()        // Store plugin.
	metricsOption := metrics.NewDefaultOption() // Metrics plugin.
	loggerOption := logger.NewDefaultOption()   // Logger plugin.
	reqlogOption := reqlog.NewDefaultOption()   // Request log plugin.
	versionOption := pversion.NewOption(        // Version plugin.
		"sundial",
		version.Version,
		version.Commit,
		version.Package,
	)

	// Enable plugins.
	cmd.EnablePlugin(metricsOption, loggerOption, reqlogOption, versionOption, dbOption)

	// Create server config.
	serverConfig := nirvana.NewConfig()

	// Configure APIs. These configurations may be changed by plugins.
	serverConfig.Configure(
		nirvana.Logger(log.DefaultLogger()), // Will be changed by logger plugin.
		nirvana.Filter(filters.Filters()...),
		nirvana.Modifier(modifiers.Modifiers()...),
		nirvana.Descriptor(apis.Descriptor()),
	)

	// Set nirvana command hooks.
	cmd.SetHook(&config.NirvanaCommandHookFunc{
		PreServeFunc: func(config *nirvana.Config, server nirvana.Server) error {
			// Output project information.
			config.Logger().Infof("Package:%s Version:%s Commit:%s", version.Package, version.Version, version.Commit)
			return nil
		},
	})

	// Start with server config.
	if err := cmd.ExecuteWithConfig(serverConfig); err != nil {
		serverConfig.Logger().Fatal(err)
	}
}
