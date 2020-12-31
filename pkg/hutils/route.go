package hutils

import (
	"net"
	"strings"
)

func RouteAdd(nifName, cidr string) error {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(ip.String(), ".0") {
		err = NifRouteHostAdd(ip.String(), nifName)
	} else {
		err = NifRouteNetAdd(ipNet.IP.String(), IPV4MaskToString(ipNet.Mask), nifName)
	}
	return err
}
