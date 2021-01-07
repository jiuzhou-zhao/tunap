package main

import (
	"github.com/jiuzhou-zhao/tunap/internal/cmd"
	"github.com/jiuzhou-zhao/tunap/pkg/minit"
)

//
// vpn route config/vpn_route.txt  --dev="以太网 2" --setup=true
//
func main() {
	minit.Init(&minit.TunAPInitConfig{ElevationURL: "http://127.0.0.1:1981"})
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
