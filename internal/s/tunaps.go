package s

import (
	"context"
	"github.com/jiuzhou-zhao/data-channel/grpc"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/jiuzhou-zhao/data-channel/dataprocessor"
	"github.com/jiuzhou-zhao/data-channel/inter"
	"github.com/jiuzhou-zhao/data-channel/tcp"
	"github.com/jiuzhou-zhao/data-channel/udp"
	"github.com/jiuzhou-zhao/data-channel/wrapper"
	udpchannel "github.com/jiuzhou-zhao/tun-channel"
	"github.com/sgostarter/i/l"
)

type TunAPServer struct {
	cfg       *Config
	logger    l.Wrapper
	udpServer *udpchannel.ChannelServer
}

func NewTunAPServer(cfg *Config, logger l.Wrapper) *TunAPServer {
	return &TunAPServer{
		cfg:    cfg,
		logger: logger.WithFields(l.StringField(l.ClsKey, "tunApiServer")),
	}
}

func (s *TunAPServer) Run() {
	var err error

	var serverChannel inter.Server

	dps := []inter.ServerDataProcessor{dataprocessor.NewServerEncryptDataProcess([]byte(s.cfg.SecKey))}

	switch strings.ToLower(s.cfg.DataChannelType) {
	case "udp", "":
		serverChannel, err = udp.NewServer(context.Background(), s.cfg.ListenAddress, nil, s.logger)
	case "tcp":
		serverChannel, err = tcp.NewServer(context.Background(), s.cfg.ListenAddress, nil, s.logger)

		dps = append(dps, dataprocessor.NewServerTCPBag())
	case "grpc":
		serverChannel, err = grpc.NewServer(context.Background(), s.cfg.ListenAddress, nil, s.logger)
	}

	if err != nil {
		s.logger.Fatal(err)
	}

	s.udpServer, err = udpchannel.NewChannelServer(context.Background(), s.logger, NewIPV4KeyParser(),
		wrapper.NewServer(serverChannel, s.logger, dps...), s.cfg.VpnVip)

	if err != nil {
		panic(err)
	}

	s.udpServer.Wait()
}

func (s *TunAPServer) GetCliInfos() (is []*CliInfo) {
	us := s.udpServer
	if us == nil {
		return
	}

	fnFormatTime := func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05")
	}

	fnFormatBytes := func(v uint64) string {
		k := uint64(1024)
		m := 1024 * k
		g := 1024 * m

		if v >= g {
			return strconv.FormatFloat(float64(v)/float64(g), 'f', 2, 64) + "GB"
		}

		if v >= m {
			return strconv.FormatFloat(float64(v)/float64(m), 'f', 2, 64) + "MB"
		}

		if v >= k {
			return strconv.FormatFloat(float64(v)/float64(k), 'f', 2, 64) + "KB"
		}

		return strconv.FormatFloat(float64(v), 'f', 2, 64) + "B"
	}

	for _, cis := range us.GetClientInfos() {
		is = append(is, &CliInfo{
			Vip: cis.VIP,
			IP:  cis.Address,
			// nolint: gosec
			VpnIPs: template.HTML(strings.Join(cis.VpnIPs, "<br>")),
			// nolint: gosec
			LanIPs:         template.HTML(strings.Join(cis.LanIPs, "<br>")),
			CreateTime:     fnFormatTime(cis.CreateTime),
			LastAccessTime: fnFormatTime(cis.LastAccessTime),
			TransBytes:     fnFormatBytes(cis.TransBytes),
		})
	}

	return
}
