package main

import (
	"strings"

	"github.com/jiuzhou-zhao/tunap/internal/c"
	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/sgostarter/i/l"
	"github.com/sgostarter/liblogrus"
)

func main() {
	logger := l.NewWrapper(liblogrus.NewLogrus())

	var cfg c.Config
	err := config.LoadConfig("c-config", &cfg)

	if err != nil {
		logger.Fatal(err)
	}

	if cfg.Env != "" && strings.ToUpper(cfg.Env) == "dev" {
		logger.GetLogger().SetLevel(l.LevelDebug)
	}

	c.NewTunAPClient(&cfg, logger).Run()
}
