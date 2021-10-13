package c

import (
	"context"
	"net"

	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"github.com/jiuzhou-zhao/tunap/pkg/minit"
	"github.com/jiuzhou-zhao/tunap/pkg/tun"
	udpchannel "github.com/jiuzhou-zhao/udp-channel"
	"github.com/jiuzhou-zhao/udp-channel/pkg"
	"github.com/sgostarter/i/logger"
)

type TunAPClient struct {
	cfg    *Config
	logger logger.Wrapper
}

func NewTunAPClient(cfg *Config, logger logger.Wrapper) *TunAPClient {
	minit.Init(&cfg.TunAPInitConfig)
	f := cfg.IsVPNTarget
	if !f {
		f = len(cfg.LanIPs) > 0
	}
	tun.ClientExtInit(f, logger)
	return &TunAPClient{
		cfg:    cfg,
		logger: logger,
	}
}

func (c *TunAPClient) dumpIPV4Package(preLog string, d []byte) {
	p := hutils.IPPacket(d)
	if p.IPver() != 4 {
		return
	}
	c.logger.Infof("%v %v -> %v", preLog, p.SrcV4().String(), p.DstV4().String())
}

func (c *TunAPClient) Run() {
	vipCidr, err := udpchannel.ToCIDR(c.cfg.Vip)
	if err != nil {
		c.logger.Fatal(err)
	}

	vip, err := udpchannel.ToIP(c.cfg.Vip)
	if err != nil {
		c.logger.Fatal(err)
	}

	tunDevice, err := tun.DeviceSetup(vipCidr)
	if err != nil {
		c.logger.Fatal(err)
	}

	vpnIPs := make([]string, 0, len(c.cfg.VpnIPs))
	for _, ip := range c.cfg.VpnIPs {
		cidr, err := udpchannel.ToCIDR(ip)
		if err != nil {
			c.logger.Fatal(err)
		}
		_ = tunDevice.RouteAdd(cidr)
		vpnIPs = append(vpnIPs, cidr)
	}

	c.logger.Infof("tun device is %v", tunDevice.Name())
	netInterface, err := net.InterfaceByName(tunDevice.Name())
	if err == nil {
		c.logger.Infof("tun device index is %v", netInterface.Index)
	}

	lanIPs := make([]string, 0, len(c.cfg.LanIPs))
	for _, ip := range c.cfg.LanIPs {
		cidr, err := udpchannel.ToCIDR(ip)
		if err != nil {
			c.logger.Fatal(err)
		}
		lanIPs = append(lanIPs, cidr)
	}

	d := &udpchannel.ChannelClientData{
		ServerAddr: c.cfg.ServerAddress,
		Key:        vip,
		VpnIPs:     vpnIPs,
		LanIPs:     lanIPs,
		Log:        c.logger,
		Crypt:      pkg.NewAESEnDecrypt(c.cfg.SecKey),
	}

	cli, err := udpchannel.NewChannelClient(context.Background(), d)
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
			c.dumpIPV4Package("READ FROM DEVICE:", d[:n])
			cli.WritePackage(d[:n])
		}
	}()

	go func() {
		for iData := range cli.ReadIncomingMsgChan() {
			if iData.Error != nil {
				c.logger.Errorf("tun device read incoming msg error: %v", iData.Error)
				continue
			}

			if iData.Data != nil {
				c.dumpIPV4Package("WRITE TO DEVICE:", iData.Data)
				_, e := tunDevice.Write(iData.Data)
				if e != nil {
					c.logger.Errorf("tun device write failed: %v", e)
				}
			}

			if len(iData.AddedForwardIPs) > 0 {
				c.logger.Info("--- AddedForwardIPs", iData.AddedForwardIPs)
				for _, cidr := range iData.AddedForwardIPs {
					_ = tunDevice.RouteAdd(cidr)
				}
			}

			if len(iData.RemovedForwardIPs) > 0 {
				c.logger.Info("--- RemovedForwardIPs", iData.RemovedForwardIPs)
				for _, cidr := range iData.RemovedForwardIPs {
					_ = tunDevice.RouteDel(cidr)
				}
			}
		}

	}()

	cli.Wait()
}
