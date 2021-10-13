// +build linux

package tun

import (
	"github.com/sgostarter/i/logger"
	"net"

	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"github.com/milosgajdos/tenus"
	"github.com/songgao/water"
)

const (
	tunDeviceInterfaceName = "tun_x_21101"
)

type LinuxTunDevice struct {
	*water.Interface
}

func (dev *LinuxTunDevice) RouteAdd(cidr string) error {
	return hutils.RouteAdd(dev.Name(), cidr)
}

func (dev *LinuxTunDevice) RouteDel(cidr string) error {
	return hutils.NifRouteHostDel(dev.Name(), cidr)
}

func DeviceSetup(localCIDR string) (TunDevice, error) {
	lIP, lNet, err := net.ParseCIDR(localCIDR)
	if err != nil {
		return nil, err
	}

	tunDev, err := water.New(water.Config{DeviceType: water.TUN, PlatformSpecificParams: water.PlatformSpecificParams{
		Name: tunDeviceInterfaceName,
	}})

	if err != nil {
		return nil, err
	}

	link, err := tenus.NewLinkFrom(tunDev.Name())
	if err != nil {
		return nil, err
	}

	err = link.SetLinkMTU(1300)
	if err != nil {
		return nil, err
	}

	err = link.SetLinkIp(lIP, lNet)
	if err != nil {
		return nil, err
	}

	err = link.SetLinkUp()
	if err != nil {
		return nil, err
	}

	err = hutils.FirewallTrustNif(tunDev.Name())
	if err != nil {
		return nil, err
	}

	return &LinuxTunDevice{
		Interface: tunDev,
	}, nil
}

func ClientExtInit(isTargetVPN bool, logger logger.Wrapper) {
	if isTargetVPN {
		err := hutils.FirewallOpenMasquerade()
		if err != nil {
			logger.Errorf("firewall open masquerade failed: %v", err)
		}
	}
}
