package command

import (
	"errors"

	"github.com/go-kit/kit/log/level"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/action"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/config"
	"github.com/urfave/cli/v2"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start integrated server",
		Flags: ServerFlags(cfg),
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			if c.IsSet("hcloud.config") {
				if err := readConfig(c.String("hcloud.config"), cfg); err != nil {
					level.Error(logger).Log(
						"msg", "Failed to read config",
						"err", err,
					)

					return err
				}
			}

			if cfg.Target.File == "" {
				level.Error(logger).Log(
					"msg", "Missing path for output.file",
				)

				return errors.New("missing path for output.file")
			}

			if c.IsSet("hcloud.token") {
				credentials := config.Credential{
					Project: "default",
					Token:   c.String("hcloud.token"),
				}

				cfg.Target.Credentials = append(
					cfg.Target.Credentials,
					credentials,
				)

				if credentials.Token == "" {
					level.Error(logger).Log(
						"msg", "Missing required hcloud.token",
					)

					return errors.New("missing required hcloud.token")
				}
			}

			if len(cfg.Target.Credentials) == 0 {
				level.Error(logger).Log(
					"msg", "Missing any credentials",
				)

				return errors.New("missing any credentials")
			}

			return action.Server(cfg, logger)
		},
	}
}

// ServerFlags defines the available server flags.
func ServerFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
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
			Name:    "hcloud.token",
			Value:   "",
			Usage:   "Access token for the HetznerCloud API",
			EnvVars: []string{"PROMETHEUS_HCLOUD_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "hcloud.config",
			Value:   "",
			Usage:   "Path to HetznerCloud configuration file",
			EnvVars: []string{"PROMETHEUS_HCLOUD_CONFIG"},
		},
	}
}
