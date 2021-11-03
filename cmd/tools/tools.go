package main

import (
	"github.com/jiuzhou-zhao/tunap/internal/cmd"
	"github.com/jiuzhou-zhao/tunap/pkg/minit"
)

//
// vpn route config/vpn_route.txt  --dev="以太网 2" --setup=true
//
/*
## linux all route
route add -host 49.234.46.74 dev ens33 gw 192.168.223.2
route add default dev tun_x_201231
# windows all route
route delete 0.0.0.0 MASK 0.0.0.0 192.168.3.1
route add 49.234.46.74 mask 255.255.255.255 192.168.3.1
route add 0.0.0.0 MASK 0.0.0.0 0.0.0.0 metric 100 if 7
*/
func main() {
	minit.Init(&minit.TunAPInitConfig{ElevationURL: "http://127.0.0.1:1981"})

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
