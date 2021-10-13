package main

import (
	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/jiuzhou-zhao/tunap/internal/s"
	"github.com/sgostarter/i/logger"
	"strings"
)

func main() {
	log := logger.NewWrapper(logger.NewCommLogger(&logger.FmtRecorder{}))
	log.GetLogger().SetLevel(logger.LevelDebug)

	var cfg s.Config
	err := config.LoadConfig("s-config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Env != "" && strings.ToUpper(cfg.Env) == "dev" {
		log.GetLogger().SetLevel(logger.LevelDebug)
	}

	s.NewTunAPServer(&cfg, log).Run()
}
