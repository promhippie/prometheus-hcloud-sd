package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/promhippie/prometheus-hcloud-sd/pkg/command"
)

func main() {
	if env := os.Getenv("PROMETHEUS_HCLOUD_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
