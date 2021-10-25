package tun

import (
	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"github.com/jiuzhou-zhao/tunap/pkg/minit"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTunDeviceSetup(t *testing.T) {
	minit.Init(&minit.TunAPInitConfig{
		ElevationURL: "http://127.0.0.1:1981",
	})

	tunDevice, err := DeviceSetup("192.168.11.2/24", "")
	assert.Nil(t, err)

	err = tunDevice.RouteAdd("192.168.11.0/24")
	assert.Nil(t, err)

	frame := make([]byte, 10000)
	for {
		n, err := tunDevice.Read(frame)
		if err != nil {
			t.Logf("err: %v", err)
			continue
		}
		t.Logf("receive %v data", n)
		ip4Package := hutils.IPPacket(frame[:n])
		t.Logf("%v -> %v", ip4Package.SrcV4().String(), ip4Package.DstV4().String())
	}

}
