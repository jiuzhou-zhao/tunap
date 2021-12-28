package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"github.com/songgao/water"
)

// nolint
func main() {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface Name: %s\n", ifce.Name())

	out, err := exec.Command("ifconfig", ifce.Name(), "100.1.0.10", "100.1.0.10", "up").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))

	out, err = exec.Command("route", "-n", "add", "-net", "20.20.20.10", "100.1.0.10").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))

	out, err = exec.Command("route", "-n", "add", "-net", "192.168.111.0", "100.1.0.10").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))

	frame := make([]byte, 2000)

	for {
		n, err := ifce.Read(frame)
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]

		dumpIPV4Package(frame)

		buf := frame
		for i := 0; i < 4; i++ {
			buf[i+12], buf[i+16] = buf[i+16], buf[i+12]
		}
		buf[20] = 0
		buf[22] = 0
		buf[23] = 0
		var checksum uint16
		for i := 20; i < n; i += 2 {
			checksum += uint16(buf[i])<<8 + uint16(buf[i+1])
		}

		checksum = ^(checksum + 4)
		buf[22] = byte(checksum >> 8)
		buf[23] = byte(checksum & ((1 << 8) - 1))

		_, err = ifce.Write(buf)
		if err != nil {
			log.Printf("error os.Write(): %v\n", err)
		}
	}
}

// nolint
func dumpIPV4Package(d []byte) {
	p := hutils.IPPacket(d)
	if p.IPver() != 4 {
		return
	}

	fmt.Printf("%v %v -> %v\n", "", p.SrcV4().String(), p.DstV4().String())
}
