package main

import (
	"os"

	"github.com/ysayonnar/zenvy/shared/logger"
)

func main() {
	log := logger.New(os.Stdout, true)

	log.Debug("hi")
	log.Info("hi")
	log.Warn("hi")
	log.Error("hi")
}
