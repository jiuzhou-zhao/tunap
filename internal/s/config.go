package s

type Config struct {
	Env              string `yaml:"Env"`              // Dev, Pro
	DataChannelType  string `yaml:"DataChannelType"`  // udp, tcp
	ListenAddress    string `yaml:"ListenAddress"`    // 服务监听的 udp 地址
	SecKey           string `yaml:"SecKey"`           // 加密key
	VpnVip           string `yaml:"VpnVip"`           // VPN模式数据发向哪个客户端
	WebListenAddress string `yaml:"WebListenAddress"` // web
}
