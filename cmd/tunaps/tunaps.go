package main

import (
	"strings"

	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/jiuzhou-zhao/tunap/internal/s"
	"github.com/sgostarter/i/logger"
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

	tunS := s.NewTunAPServer(&cfg, log)
	if cfg.WebListenAddress != "" {
		go func() {
			s.RunWeb(cfg.WebListenAddress, tunS)
		}()
	}

	tunS.Run()
}
