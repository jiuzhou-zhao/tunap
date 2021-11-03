package main

import (
	"strings"

	"github.com/jiuzhou-zhao/tunap/internal/c"
	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/sgostarter/i/logger"
)

func main() {
	rLog := logger.NewCommLogger(&logger.FmtRecorder{})
	rLog.SetLevel(logger.LevelDebug)
	log := logger.NewWrapper(rLog)

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
