// +build linux

package hutils

import (
	"os/exec"
	"strings"
)

func executeEx(name string, args []string) (o string, err error) {
	out, err := exec.Command(name, args...).CombinedOutput()
	if out != nil {
		o = string(out)
	}
	return
}

func execute(name string, args []string) error {
	_, err := executeEx(name, args)
	return err
}

func FirewallTrustNif(nifName string) error {
	// firewall-cmd --zone=trusted --add-interface=tun0
	o, _ := executeEx("firewall-cmd", []string{
		"--zone=trusted",
		"--query-interface=" + nifName,
	})
	if strings.HasPrefix(o, "yes") {
		return nil
	}
	_ = execute("firewall-cmd", []string{
		"--remove-interface=" + nifName,
	})

	return execute("firewall-cmd", []string{
		"--zone=trusted",
		"--add-interface=" + nifName,
	})
}

func NifRouteNetAdd(ipNet, ipNetmask, dev string) error {
	return execute("route", []string{
		"add",
		"-net",
		ipNet,
		"netmask",
		ipNetmask,
		"dev",
		dev,
	})
}

func NifRouteHostAdd(host, dev string) error {
	return execute("route", []string{
		"add",
		"-host",
		host,
		"dev",
		dev,
	})
}
