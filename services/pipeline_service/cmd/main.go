package main

import (
	log2 "log"
	"os"

	"github.com/ysayonnar/zenvy/services/pipelines-service/internal/config"
	"github.com/ysayonnar/zenvy/shared/env"
	"github.com/ysayonnar/zenvy/shared/logger"
)

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log2.Fatalf("Config not parsed: %s", err)
	}

	log := logger.New(os.Stdout, cfg.IsDebug)

	log.Info("config parsed")

	log.Debug("hi")
	log.Info("hi")
	log.Warn("hi")
	log.Error("hi")

	for {
	}
}
