// +build darwin

package mos

func NifRouteNetAdd(ipNet, ipNetmask, dev string) error {
	return nil
}

func NifRouteNetDel(ipNet, ipNetmask, dev string) error {
	return nil
}

func NifRouteHostAdd(host, dev string) error {
	return nil
}

func NifRouteHostDel(host, dev string) error {
	return nil
}
