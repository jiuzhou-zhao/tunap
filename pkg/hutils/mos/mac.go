//go:build darwin
// +build darwin

package mos

import "os/exec"

func NifSetIPAddress(ifName, ip, mask string) error {
	_, err := exec.Command("ifconfig", ifName, ip, ip, "up").CombinedOutput()

	return err
}

func NifRouteNetAdd(ipNet, ipNetmask, dev string) error {
	args := []string{
		"-n",
		"add",
		"-net",
		ipNet,
	}
	if ipNetmask != "" {
		args = append(args, "-netmask")
		args = append(args, ipNetmask)
	}

	args = append(args, dev)
	_, err := exec.Command("route", args...).CombinedOutput()

	return err
}

func NifRouteNetDel(ipNet, ipNetmask, dev string) error {
	// route -v delete -net
	_, err := exec.Command("route", "-v", "delete", "-net", ipNet).CombinedOutput()

	return err
}

func NifRouteHostAdd(host, dev string) error {
	return NifRouteNetAdd(host, "", dev)
}

func NifRouteHostDel(host, dev string) error {
	return NifRouteNetDel(host, "", dev)
}
