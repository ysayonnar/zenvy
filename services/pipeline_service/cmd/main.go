package main

import (
	"context"
	log2 "log"
	"os"
	"time"

	"github.com/ysayonnar/zenvy/services/pipelines-service/db"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pg, err := db.Connect(ctx, "postgres", cfg.Postgres.DSN())
	if err != nil {
		log.Error("postgres connect error", logger.Err(err))
		os.Exit(1)
	}
	defer pg.Close()

	log.Info("postgres connected")

	for {
	}
}
