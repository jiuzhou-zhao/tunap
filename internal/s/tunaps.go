package s

import (
	"context"
	udpchannel "github.com/jiuzhou-zhao/udp-channel"
	"github.com/jiuzhou-zhao/udp-channel/pkg"
	"github.com/sgostarter/i/logger"
)

type TunAPServer struct {
	cfg    *Config
	logger logger.Wrapper
}

func NewTunAPServer(cfg *Config, logger logger.Wrapper) *TunAPServer {
	return &TunAPServer{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *TunAPServer) Run() {
	udpServer, err := udpchannel.NewChannelServer(context.Background(), s.cfg.ListenAddress, s.logger, NewIPV4KeyParser(),
		pkg.NewAESEnDecrypt(s.cfg.SecKey), s.cfg.VpnVip)
	if err != nil {
		panic(err)
	}
	udpServer.Wait()
}
