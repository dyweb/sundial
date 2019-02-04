package plugin

import (
	"context"

	"github.com/caicloud/nirvana"
	def "github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/service"
	"github.com/dyweb/sundial/pkg/store/rdb"
	rdbstore "github.com/dyweb/sundial/pkg/store/rdb/datastore"
	"github.com/dyweb/sundial/pkg/store/tsdb"
	influx "github.com/dyweb/sundial/pkg/store/tsdb/influx"
)

const (
	// ExternalConfigName is the config name for the plugin.
	ExternalConfigName = "db"
)

func init() {
	nirvana.RegisterConfigInstaller(&DBInstaller{})
}

// DBInstaller is the installer for DB plugin in nirvana.
type DBInstaller struct{}

// Name is the external config name.
func (i *DBInstaller) Name() string {
	return ExternalConfigName
}

// Install installs stuffs before server starting.
func (i *DBInstaller) Install(builder service.Builder, cfg *nirvana.Config) error {
	var err error
	wrapper(cfg, func(c *Option) {
		err = builder.AddDescriptor(def.Descriptor{
			Path: "/",
			Middlewares: []def.Middleware{
				func(ctx context.Context, next def.Chain) error {
					s := rdbstore.New(c.RDBDriver, c.RDBSource)
					ts := influx.New(c.TSDBSource, c.TSDBUsername, c.TSDBPassword, c.TSDBName)
					newCtx := tsdb.ToContext(rdb.ToContext(ctx, s), ts)
					return next.Continue(newCtx)
				},
			},
		})
	})
	return err
}

// Uninstall uninstalls stuffs after server terminating.
func (i *DBInstaller) Uninstall(builder service.Builder, cfg *nirvana.Config) error {
	return nil
}

// Option is the config for DB related options.
type Option struct {
	RDBDriver    string `desc:"Relational database system driver (mysql, sqlite3, mssql)"`
	RDBSource    string `desc:"Relational database system source"`
	TSDBDriver   string `desc:"Time series database system driver (influx)"`
	TSDBSource   string `desc:"Time series database system source"`
	TSDBUsername string `desc:"Username of the time series dababase system"`
	TSDBPassword string `desc:"Password of the time series dababase system"`
	TSDBName     string `desc:"Name of the database in the time series dababase system"`
}

// NewDefaultOption creates default option.
func NewDefaultOption() *Option {
	return &Option{
		RDBDriver:    "sqlite3",
		RDBSource:    "sundial.sqlite",
		TSDBDriver:   "influx",
		TSDBSource:   "http://localhost:8086",
		TSDBUsername: "user",
		TSDBPassword: "password",
		TSDBName:     "sundial",
	}
}

// Name returns plugin name.
func (p *Option) Name() string {
	return ExternalConfigName
}

// Configure configures nirvana config via current options.
func (p *Option) Configure(cfg *nirvana.Config) error {
	cfg.Configure(
		RDBDriver(p.RDBDriver),
		RDBSource(p.RDBSource),
		TSDBDriver(p.TSDBDriver),
		TSDBSource(p.TSDBSource),
		TSDBUsername(p.TSDBUsername),
		TSDBPassword(p.TSDBPassword),
		TSDBName(p.TSDBName),
	)
	return nil
}

// Disable returns a configurer to disable reqlog.
func Disable() nirvana.Configurer {
	return func(c *nirvana.Config) error {
		c.Set(ExternalConfigName, nil)
		return nil
	}
}

func RDBDriver(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.RDBDriver = s
		})
		return nil
	}
}

func RDBSource(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.RDBSource = s
		})
		return nil
	}
}

func TSDBDriver(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.TSDBDriver = s
		})
		return nil
	}
}

func TSDBSource(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.TSDBSource = s
		})
		return nil
	}
}

func TSDBUsername(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.TSDBUsername = s
		})
		return nil
	}
}

func TSDBPassword(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.TSDBPassword = s
		})
		return nil
	}
}

func TSDBName(s string) nirvana.Configurer {
	return func(c *nirvana.Config) error {
		wrapper(c, func(c *Option) {
			c.TSDBName = s
		})
		return nil
	}
}

func wrapper(c *nirvana.Config, f func(c *Option)) {
	conf := c.Config(ExternalConfigName)
	var cfg *Option
	if conf == nil {
		// Default config.
		cfg = NewDefaultOption()
	} else {
		// Panic if config type is wrong.
		cfg = conf.(*Option)
	}
	f(cfg)
}
