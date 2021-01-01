package c

import (
	"context"
	"net"

	"github.com/jiuzhou-zhao/tunap/pkg/logrus-logger"
	"github.com/jiuzhou-zhao/tunap/pkg/minit"
	"github.com/jiuzhou-zhao/tunap/pkg/tun"
	"github.com/jiuzhou-zhao/udp-channel"
	"github.com/jiuzhou-zhao/udp-channel/pkg"
)

type TunAPClient struct {
	cfg    *Config
	logger *logrus_logger.Logger
}

func NewTunAPClient(cfg *Config, logger *logrus_logger.Logger) *TunAPClient {
	minit.Init(&cfg.TunAPInitConfig)
	tun.TunClientExtInit(cfg.IsVPNTarget, logger)
	return &TunAPClient{
		cfg:    cfg,
		logger: logger,
	}
}

func (c *TunAPClient) Run() {
	ip, _, err := net.ParseCIDR(c.cfg.VipCIDR)
	if err != nil {
		c.logger.Fatal(err)
	}

	tunDevice, err := tun.TunDeviceSetup(c.cfg.VipCIDR)
	if err != nil {
		c.logger.Fatal(err)
	}

	cli, err := udp_channel.NewChannelClient(context.Background(), c.cfg.ServerAddress, ip.To4().String(),
		c.logger, pkg.NewAESEnDecrypt(c.cfg.SecKey))
	if err != nil {
		c.logger.Fatal(err)
	}

	go func() {
		for {
			d := make([]byte, 40960)
			n, e := tunDevice.Read(d)
			if e != nil {
				c.logger.Errorf("tun device read failed: %v", e)
				continue
			}
			cli.WritePackage(d[:n])
		}
	}()

	go func() {
		for d := range cli.ReadPackageChan() {
			_, e := tunDevice.Write(d)
			if e != nil {
				c.logger.Errorf("tun device write failed: %v", e)
				continue
			}
		}

	}()

	cli.Wait()
}
