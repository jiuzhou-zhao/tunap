package s

import (
	"context"
	udpchannel "github.com/jiuzhou-zhao/udp-channel"
	"github.com/jiuzhou-zhao/udp-channel/pkg"
	"github.com/sgostarter/i/logger"
	"html/template"
	"strings"
)

type TunAPServer struct {
	cfg       *Config
	logger    logger.Wrapper
	udpServer *udpchannel.ChannelServer
}

func NewTunAPServer(cfg *Config, logger logger.Wrapper) *TunAPServer {
	return &TunAPServer{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *TunAPServer) Run() {
	var err error
	s.udpServer, err = udpchannel.NewChannelServer(context.Background(), s.cfg.ListenAddress, s.logger, NewIPV4KeyParser(),
		pkg.NewAESEnDecrypt(s.cfg.SecKey), s.cfg.VpnVip)
	if err != nil {
		panic(err)
	}
	s.udpServer.Wait()
}

func (s *TunAPServer) GetCliInfos() (is []*CliInfo) {
	us := s.udpServer
	if us == nil {
		return
	}

	for _, cis := range us.GetClientInfos() {
		is = append(is, &CliInfo{
			Vip:    cis.VIP,
			Ip:     cis.Address,
			VpnIPs: template.HTML(strings.Join(cis.VpnIPs, "<br>")),
			LanIPs: template.HTML(strings.Join(cis.LanIPs, "<br>")),
		})
	}

	return
}
