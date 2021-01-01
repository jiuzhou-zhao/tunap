package main

import (
	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/jiuzhou-zhao/tunap/internal/s"
	"github.com/jiuzhou-zhao/tunap/pkg/logrus-logger"
)

func main() {
	logger := logrus_logger.NewLogger(nil)

	var cfg s.Config
	err := config.LoadConfig("s-config", &cfg)
	if err != nil {
		logger.Fatal(err)
	}
	s.NewTunAPServer(&cfg, logger).Run()
}
