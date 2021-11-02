package c

import "github.com/jiuzhou-zhao/tunap/pkg/minit"

type Config struct {
	minit.TunAPInitConfig `yaml:",inline"`
	DataChannelType       string   `yaml:"DataChannelType"` // udp, tcp
	Env                   string   `yaml:"Env"`             // Dev, Pro
	ServerAddress         string   `yaml:"ServerAddress"`   // tunaps 服务的 udp 地址
	SecKey                string   `yaml:"SecKey"`          // 加密key
	Vip                   string   `yaml:"Vip"`             // 本client的IP地址
	IsVPNTarget           bool     `yaml:"IsVPNTarget"`     // 是否为 VPN Target
	VpnIPs                []string `yaml:"VpnIPs"`          // VPN IP 列表
	LanIPs                []string `yaml:"LanIPs"`          // 局域网 IP 列表
	NifName               string   `yaml:"NifName"`         // 网卡名字
}
