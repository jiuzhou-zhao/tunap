package cmd

import (
	"bufio"
	"github.com/jiuzhou-zhao/tunap/pkg/hutils/mos"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var dev string
var setup bool
var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "vpn",
	Long:  "vpn",
	Args:  cobra.MinimumNArgs(1),
}

func routesFromFile(file string) (routes []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()

	routes = make([]string, 0, 100)
	r := bufio.NewReader(f)

	var line []byte
	for {
		line, _, err = r.ReadLine()
		if err == io.EOF {
			err = nil
			break
		}
		lineS := string(line)
		lineS = strings.Trim(lineS, "\r\n \t")
		if lineS == "" || lineS[0] == '#' {
			continue
		}
		ls := strings.Split(lineS, "=")
		if len(ls) != 2 {
			continue
		}
		ls[0] = strings.Trim(ls[0], "\r\n \t")
		ls[1] = strings.Trim(ls[1], "\r\n \t")
		if ls[0] != "route" {
			continue
		}
		routes = append(routes, ls[1])
	}
	return
}

var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "setup/cleanup vpn route",
	Long:  `setup/cleanup vpn route`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			return
		}
		routes, err := routesFromFile(args[0])
		if err != nil {
			logrus.Errorf("read routes from file [%v] failed: %v", args[0], err)
			return
		}
		if setup {
			mos.SetRoutes(routes, dev)
		} else {
			mos.UnsetRoutes(routes, dev)
		}
	},
}

func init() {
	rootCmd.AddCommand(vpnCmd)
	vpnCmd.AddCommand(routeCmd)

	routeCmd.Flags().StringVar(&dev, "dev", "", "tun/tap device name")
	routeCmd.Flags().BoolVar(&setup, "setup", true, "setup or cleanup vpn routes")
}
