package command

import (
	"os"

	"github.com/promhippie/prometheus-hcloud-sd/pkg/config"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/version"
	"github.com/urfave/cli/v3"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.App{
		Name:    "prometheus-hcloud-sd",
		Version: version.String,
		Usage:   "Prometheus HetznerCloud SD",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: RootFlags(cfg),
		Commands: []*cli.Command{
			Health(cfg),
			Server(cfg),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	return app.Run(os.Args)
}

// RootFlags defines the available root flags.
func RootFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "info",
			Usage:       "Only log messages with given severity",
			EnvVars:     []string{"PROMETHEUS_HCLOUD_LOG_LEVEL"},
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log.pretty",
			Value:       false,
			Usage:       "Enable pretty messages for logging",
			EnvVars:     []string{"PROMETHEUS_HCLOUD_LOG_PRETTY"},
			Destination: &cfg.Logs.Pretty,
		},
	}
}
