package minit

type TunAPInitConfig struct {
	ElevationURL string `json:"elevation_url"`
}

var (
	Cfg TunAPInitConfig
)

func Init(cfg *TunAPInitConfig) {
	Cfg = *cfg
}
