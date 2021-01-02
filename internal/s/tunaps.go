package s

import (
	"context"
	"github.com/jiuzhou-zhao/tunap/pkg/logrus-logger"
	"github.com/jiuzhou-zhao/udp-channel"
	"github.com/jiuzhou-zhao/udp-channel/pkg"
)

type TunAPServer struct {
	cfg    *Config
	logger *logrus_logger.Logger
}

func NewTunAPServer(cfg *Config, logger *logrus_logger.Logger) *TunAPServer {
	return &TunAPServer{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *TunAPServer) Run() {
	udpServer, err := udp_channel.NewChannelServer(context.Background(), s.cfg.ListenAddress, s.logger, NewIPV4KeyParser(),
		pkg.NewAESEnDecrypt(s.cfg.SecKey), s.cfg.VpnVip)
	if err != nil {
		panic(err)
	}
	udpServer.Wait()
}
