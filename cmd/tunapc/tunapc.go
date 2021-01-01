package main

import (
	"github.com/jiuzhou-zhao/tunap/internal/c"
	"github.com/jiuzhou-zhao/tunap/internal/config"
	"github.com/jiuzhou-zhao/tunap/pkg/logrus-logger"
)

func main() {
	logger := logrus_logger.NewLogger(nil)

	var cfg c.Config
	err := config.LoadConfig("c-config", &cfg)
	if err != nil {
		logger.Fatal(err)
	}
	c.NewTunAPClient(&cfg, logger).Run()
}
