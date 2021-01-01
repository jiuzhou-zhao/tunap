package c

import "github.com/jiuzhou-zhao/tunap/pkg/minit"

type Config struct {
	minit.TunAPInitConfig `yaml:",inline"`
	ServerAddress         string `yaml:"ServerAddress"` // tunaps 服务的 udp 地址
	SecKey                string `yaml:"SecKey"`        // 加密key
	VipCIDR               string `yaml:"VipCIDR"`       // 本client的CIDR地址 当前只支持 IPV4 X.X.X.X/24 的格式
	IsVPNTarget           bool   `yaml:"IsVPNTarget"`   // 是否为 VPN Target
}
