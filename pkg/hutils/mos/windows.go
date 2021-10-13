// +build windows

package mos

import (
	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"net"
	"strconv"
)

func NifSetIPAddress(ifName, ip, mask string) error {
	return hutils.ElevationExecute("netsh", []string{
		"interface",
		"ip",
		"set",
		"address",
		"name=" + ifName,
		"source=static",
		"addr=" + ip,
		"mask=" + mask,
		"gateway=none",
	})
}

func NifRouteNetAdd(ipNet, ipNetmask, dev string) error {
	netInterface, err := net.InterfaceByName(dev)
	if err != nil {
		return err
	}
	return hutils.ElevationExecute("route", []string{
		"add",
		ipNet,
		"mask",
		ipNetmask,
		"0.0.0.0",
		"if",
		strconv.Itoa(netInterface.Index),
	})
}

func NifRouteNetDel(ipNet, ipNetmask, dev string) error {
	//netInterface, err := net.InterfaceByName(dev)
	//if err != nil {
	//	return err
	//}
	return hutils.ElevationExecute("route", []string{
		"delete",
		ipNet,
		"mask",
		ipNetmask,
		//"0.0.0.0",
		//"if",
		//strconv.Itoa(netInterface.Index),
	})
}

func NifRouteHostAdd(host, dev string) error {
	netInterface, err := net.InterfaceByName(dev)
	if err != nil {
		return err
	}
	return hutils.ElevationExecute("route", []string{
		"add",
		host,
		"0.0.0.0",
		"if",
		strconv.Itoa(netInterface.Index),
	})
}

func NifRouteHostDel(host, dev string) error {
	// netInterface, err := net.InterfaceByName(dev)
	// if err != nil {
	//	return err
	//}
	return hutils.ElevationExecute("route", []string{
		"delete",
		host,
		"0.0.0.0", /*
			"if",
			strconv.Itoa(netInterface.Index),*/
	})
}
