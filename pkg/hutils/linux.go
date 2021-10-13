// +build linux

package hutils

import (
	"github.com/golang/glog"
	"os/exec"
	"strings"
)

const (
	firewallCmd = "firewall-cmd"
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

// 信任某个NIF，则连接该NIF的端口不经过防火墙
func FirewallTrustNif(nifName string) error {
	// firewall-cmd --zone=trusted --add-interface=tun0
	o, _ := executeEx(firewallCmd, []string{
		"--zone=trusted",
		"--query-interface=" + nifName,
	})
	if strings.HasPrefix(o, "yes") {
		return nil
	}
	_ = execute(firewallCmd, []string{
		"--remove-interface=" + nifName,
	})

	return execute(firewallCmd, []string{
		"--zone=trusted",
		"--add-interface=" + nifName,
	})
}

func FirewallQueryMasquerade() bool {
	o, _ := executeEx(firewallCmd, []string{
		"--query-masquerade",
	})
	return strings.HasPrefix(o, "yes")
}

// 防火墙的masquerade功能进行地址伪装（NAT），私网访问公网或公网访问私网都需要开启此功能来进行地址转换，否则无法正常互访
func FirewallOpenMasquerade() error {
	if FirewallQueryMasquerade() {
		return nil
	}
	_, err := executeEx(firewallCmd, []string{
		"--add-masquerade",
	})
	if err != nil {
		glog.Errorf("FwTrustInterface failed: %v", err)
		return err
	}
	return nil
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

func NifRouteNetDel(ipNet, ipNetmask, dev string) error {
	return execute("route", []string{
		"del",
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

func NifRouteHostDel(host, dev string) error {
	return execute("route", []string{
		"del",
		"-host",
		host,
		"dev",
		dev,
	})
}
