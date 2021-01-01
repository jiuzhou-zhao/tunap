package s

type Config struct {
	ListenAddress string `yaml:"ListenAddress"` // 服务监听的 udp 地址
	SecKey        string `yaml:"SecKey"`        // 加密key
	VpnVip        string `yaml:"VpnVip"`        // VPN模式数据发向哪个客户端
}
