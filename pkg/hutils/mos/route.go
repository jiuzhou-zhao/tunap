package mos

import (
	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
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
		err = NifRouteNetAdd(ipNet.IP.String(), hutils.IPV4MaskToString(ipNet.Mask), nifName)
	}
	return err
}

func RouteDel(nifName, cidr string) error {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(ip.String(), ".0") {
		err = NifRouteHostDel(ip.String(), nifName)
	} else {
		err = NifRouteNetDel(ipNet.IP.String(), hutils.IPV4MaskToString(ipNet.Mask), nifName)
	}
	return err
}

func RouteAddByNetMask(nifName, lpNet, lpMask string) (err error) {
	if !strings.HasSuffix(lpNet, ".0") {
		err = NifRouteHostAdd(lpNet, nifName)
	} else {
		err = NifRouteNetAdd(lpNet, lpMask, nifName)
	}
	return
}

func RouteDelByNetMask(nifName, lpNet, lpMask string) (err error) {
	if !strings.HasSuffix(lpNet, ".0") {
		err = NifRouteHostDel(lpNet, nifName)
	} else {
		err = NifRouteNetDel(lpNet, lpMask, nifName)
	}
	return
}

func parseNifRouteItem(route string) (net string, mask string, ok bool) {
	rs := strings.Split(route, "/")
	if len(rs) != 2 {
		return
	}
	if len(strings.Split(rs[0], ".")) != 4 || len(strings.Split(rs[1], ".")) != 4 {
		return
	}
	net = rs[0]
	mask = rs[1]
	ok = true
	return
}

func SetRoutes(routes []string, dev string) []string {
	effectedRoutes := make([]string, 0, len(routes))
	for _, route := range routes {
		n, mask, ok := parseNifRouteItem(route)
		if !ok {
			continue
		}
		err := RouteAddByNetMask(dev, n, mask)
		if err != nil {
			continue
		}
		effectedRoutes = append(effectedRoutes, route)
	}
	return effectedRoutes
}

func UnsetRoutes(routes []string, dev string) {
	for _, route := range routes {
		n, mask, ok := parseNifRouteItem(route)
		if !ok {
			continue
		}
		err := RouteDelByNetMask(dev, n, mask)
		if err != nil {
			continue
		}
	}
}
