package main

import (
	"github.com/jiuzhou-zhao/tunap/internal/c"
	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/sgostarter/i/logger"
	"strings"
)

func main() {
	log := logger.NewWrapper(logger.NewCommLogger(&logger.FmtRecorder{}))

	var cfg c.Config
	err := config.LoadConfig("c-config", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Env != "" && strings.ToUpper(cfg.Env) == "dev" {
		log.GetLogger().SetLevel(logger.LevelDebug)
	}

	c.NewTunAPClient(&cfg, log).Run()
}
