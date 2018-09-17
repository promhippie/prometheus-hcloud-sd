package main

import (
	"errors"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/action"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/config"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

var (
	// ErrMissingOutputFile defines the error if output.file is empty.
	ErrMissingOutputFile = errors.New("Missing path for output.file")

	// ErrMissingHcloudToken defines the error if hcloud.token is empty.
	ErrMissingHcloudToken = errors.New("Missing required hcloud.token")
)

func main() {
	cfg := config.Load()

	if env := os.Getenv("PROMETHEUS_HCLOUD_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:    "prometheus-hcloud-sd",
		Version: version.Version,
		Usage:   "Prometheus HetznerCloud SD",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: []cli.Flag{
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
		},
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "start integrated server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "web.address",
						Value:       "0.0.0.0:9000",
						Usage:       "Address to bind the metrics server",
						EnvVars:     []string{"PROMETHEUS_HCLOUD_WEB_ADDRESS"},
						Destination: &cfg.Server.Addr,
					},
					&cli.StringFlag{
						Name:        "web.path",
						Value:       "/metrics",
						Usage:       "Path to bind the metrics server",
						EnvVars:     []string{"PROMETHEUS_HCLOUD_WEB_PATH"},
						Destination: &cfg.Server.Path,
					},
					&cli.StringFlag{
						Name:        "output.file",
						Value:       "/etc/prometheus/hcloud.json",
						Usage:       "Path to write the file_sd config",
						EnvVars:     []string{"PROMETHEUS_HCLOUD_OUTPUT_FILE"},
						Destination: &cfg.Target.File,
					},
					&cli.IntFlag{
						Name:        "output.refresh",
						Value:       30,
						Usage:       "Discovery refresh interval in seconds",
						EnvVars:     []string{"PROMETHEUS_HCLOUD_OUTPUT_REFRESH"},
						Destination: &cfg.Target.Refresh,
					},
					&cli.StringFlag{
						Name:        "hcloud.token",
						Value:       "",
						Usage:       "Access token for the HetznerCloud API",
						EnvVars:     []string{"PROMETHEUS_HCLOUD_TOKEN"},
						Destination: &cfg.Target.Token,
					},
				},
				Action: func(c *cli.Context) error {
					logger := setupLogger(cfg)

					if cfg.Target.File == "" {
						level.Error(logger).Log(
							"msg", ErrMissingOutputFile,
						)

						return ErrMissingOutputFile
					}

					if cfg.Target.Token == "" {
						level.Error(logger).Log(
							"msg", ErrMissingHcloudToken,
						)

						return ErrMissingHcloudToken
					}

					return action.Server(cfg, logger)
				},
			},
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

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
