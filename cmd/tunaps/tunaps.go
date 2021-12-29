package main

import (
	"strings"

	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/jiuzhou-zhao/tunap/internal/s"
	"github.com/sgostarter/i/l"
	"github.com/sgostarter/liblogrus"
)

func main() {
	logger := l.NewWrapper(liblogrus.NewLogrus())
	logger.GetLogger().SetLevel(l.LevelWarn)

	var cfg s.Config
	err := config.LoadConfig("s-config", &cfg)

	if err != nil {
		logger.Fatal(err)
	}

	if cfg.Env != "" && strings.ToUpper(cfg.Env) == "dev" {
		logger.GetLogger().SetLevel(l.LevelDebug)
	}

	tunS := s.NewTunAPServer(&cfg, logger)

	if cfg.WebListenAddress != "" {
		go func() {
			s.RunWeb(cfg.WebListenAddress, tunS)
		}()
	}

	tunS.Run()
}
