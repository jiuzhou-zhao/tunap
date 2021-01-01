package minit

type TunAPInitConfig struct {
	// 暂时windows才需要，linux系统直接root运行
	ElevationURL string `yaml:"ElevationURL"`
}

var (
	Cfg TunAPInitConfig
)

func Init(cfg *TunAPInitConfig) {
	Cfg = *cfg
}
