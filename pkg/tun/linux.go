//go:build linux
// +build linux

package tun

import (
	"net"

	"github.com/jiuzhou-zhao/tunap/pkg/hutils/mos"
	"github.com/milosgajdos/tenus"
	"github.com/sgostarter/i/l"
	"github.com/songgao/water"
)

const (
	tunDeviceInterfaceName = "tun_x_21101"
)

type LinuxTunDevice struct {
	*water.Interface
}

func (dev *LinuxTunDevice) RouteAdd(cidr string) error {
	return mos.RouteAdd(dev.Name(), cidr)
}

func (dev *LinuxTunDevice) RouteDel(cidr string) error {
	return mos.NifRouteHostDel(dev.Name(), cidr)
}

func DeviceSetup(localCIDR, deviceName string) (Device, error) {
	lIP, lNet, err := net.ParseCIDR(localCIDR)
	if err != nil {
		return nil, err
	}

	if deviceName == "" {
		deviceName = tunDeviceInterfaceName
	}

	tunDev, err := water.New(water.Config{DeviceType: water.TUN, PlatformSpecificParams: water.PlatformSpecificParams{
		Name: deviceName,
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

	err = mos.FirewallTrustNif(tunDev.Name())
	if err != nil {
		return nil, err
	}

	return &LinuxTunDevice{
		Interface: tunDev,
	}, nil
}

func ClientExtInit(isTargetVPN bool, logger l.Wrapper) {
	if isTargetVPN {
		err := mos.FirewallOpenMasquerade()
		if err != nil {
			logger.Errorf("firewall open masquerade failed: %v", err)
		}
	}
}
