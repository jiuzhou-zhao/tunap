//go:build darwin
// +build darwin

package tun

import (
	"net"

	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"github.com/jiuzhou-zhao/tunap/pkg/hutils/mos"
	"github.com/sgostarter/i/logger"
	"github.com/songgao/water"
)

type MacTunDevice struct {
	*water.Interface
	nifIP string
}

func (dev *MacTunDevice) RouteAdd(cidr string) error {
	return mos.RouteAdd(dev.Name(), cidr)
}

func (dev *MacTunDevice) RouteDel(cidr string) error {
	return mos.NifRouteHostDel(dev.Name(), cidr)
}

func (dev *MacTunDevice) Name() string {
	return dev.nifIP
}

func DeviceSetup(localCIDR, _deviceName string) (Device, error) {
	lIP, lNet, err := net.ParseCIDR(localCIDR)
	if err != nil {
		return nil, err
	}

	tunDev, err := water.New(water.Config{DeviceType: water.TUN, PlatformSpecificParams: water.PlatformSpecificParams{}})

	if err != nil {
		return nil, err
	}

	err = mos.NifSetIPAddress(tunDev.Name(), lIP.String(), hutils.IPV4MaskToString(lNet.Mask))
	if err != nil {
		return nil, err
	}

	_ = mos.NifRouteNetAdd(lNet.String(), "", lIP.String())

	return &MacTunDevice{
		Interface: tunDev,
		nifIP:     lIP.String(),
	}, nil
}

func ClientExtInit(isTargetVPN bool, logger logger.Wrapper) {

}
