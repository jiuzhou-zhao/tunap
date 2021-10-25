// +build windows

package tun

import (
	"errors"
	"github.com/jiuzhou-zhao/tunap/pkg/hutils/mos"
	"net"
	"sync"

	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	"github.com/jiuzhou-zhao/udp-channel/pkg"
	"github.com/mdlayher/arp"
	"github.com/mdlayher/ethernet"
	"github.com/songgao/water"
)

type WinTunDevice struct {
	sync.RWMutex
	ifc        *water.Interface
	ip         net.IP
	ipNet      *net.IPNet
	realMacMap map[string]net.HardwareAddr
}

func (dev *WinTunDevice) fakeHardwareAddr(ip net.IP) net.HardwareAddr {
	ip = ip.To4()
	return net.HardwareAddr{0xf0, 0x18, ip[0], ip[1], ip[2], ip[3]}
}

func (dev *WinTunDevice) recordHardwareAddr(ip net.IP, addr net.HardwareAddr) {
	dev.Lock()
	defer dev.Unlock()
	dev.realMacMap[ip.String()] = addr
}

func (dev *WinTunDevice) getHardwareAddr(ip net.IP) (net.HardwareAddr, error) {
	dev.RLock()
	defer dev.RUnlock()

	if addr, ok := dev.realMacMap[ip.String()]; ok {
		return addr, nil
	}
	return dev.fakeHardwareAddr(ip), nil
	// return nil, fmt.Errorf("no mac address record for %v", ip.String())
}

func (dev *WinTunDevice) writeEthernetFrame(data []byte, destHardwareAddr net.HardwareAddr, srcHardwareAddr net.HardwareAddr, etherType ethernet.EtherType) error {
	frame := &ethernet.Frame{
		Destination: destHardwareAddr,
		Source:      srcHardwareAddr,
		EtherType:   etherType,
		Payload:     data,
	}
	fb, err := frame.MarshalBinary()
	if err != nil {
		return err
	}
	n, err := dev.ifc.Write(fb)
	if err != nil || n != len(fb) {
		return err
	}
	return nil
}

func (dev *WinTunDevice) replyArp(arpFrame *arp.Packet) {
	arpReply, err := arp.NewPacket(arp.OperationReply, dev.fakeHardwareAddr(arpFrame.TargetIP), arpFrame.TargetIP,
		arpFrame.SenderHardwareAddr, arpFrame.SenderIP)
	if err != nil {
		return
	}
	arpData, err := arpReply.MarshalBinary()
	if err != nil {
		return
	}
	err = dev.writeEthernetFrame(arpData, arpFrame.SenderHardwareAddr, dev.fakeHardwareAddr(arpFrame.TargetIP), ethernet.EtherTypeARP)
	if err != nil {
		return
	}
}

func (dev *WinTunDevice) Read(p []byte) (n int, err error) {
read:
	frame := make([]byte, 10000)
	n, err = dev.ifc.Read(frame)
	if err != nil {
		return
	}
	frame = frame[:n]
	arpFrame, ethernetFrame, err := hutils.ParseEthernetFrame(frame)
	if err != nil {
		return
	}
	if arpFrame != nil {
		// arp
		if arpFrame.Operation == arp.OperationRequest {
			dev.recordHardwareAddr(arpFrame.SenderIP, arpFrame.SenderHardwareAddr)
		}
		if dev.ip.String() != arpFrame.TargetIP.String() { //  && (dev.ipNet.Contains(arpFrame.TargetIP))
			dev.replyArp(arpFrame)
		}
		//
		goto read
	}

	if ethernetFrame.EtherType == ethernet.EtherTypeIPv4 {
		n = copy(p, ethernetFrame.Payload)
		return
	}

	goto read
}

func (dev *WinTunDevice) Write(p []byte) (n int, err error) {
	ipPackage := hutils.IPPacket(p)
	if ipPackage.IPver() != 4 {
		err = errors.New("not ip v4 package")
		return
	}
	dstAddr, err := dev.getHardwareAddr(ipPackage.DstV4())
	if err != nil {
		return
	}
	err = dev.writeEthernetFrame(p, dstAddr, dev.fakeHardwareAddr(ipPackage.SrcV4()), ethernet.EtherTypeIPv4)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (dev *WinTunDevice) Close() error {
	return dev.ifc.Close()
}

func (dev *WinTunDevice) Name() string {
	return dev.ifc.Name()
}

func (dev *WinTunDevice) RouteAdd(cidr string) error {
	return mos.RouteAdd(dev.Name(), cidr)
}

func TunDeviceSetup(localCIDR, _ string) (TunDevice, error) {
	lIP, lNet, err := net.ParseCIDR(localCIDR)
	if err != nil {
		return nil, err
	}

	tunDev, err := water.New(water.Config{DeviceType: water.TAP, PlatformSpecificParams: water.PlatformSpecificParams{
		ComponentID: "tap0901",
		Network:     localCIDR,
	}})

	if err != nil {
		return nil, err
	}

	err = hutils.NifSetIPAddress(tunDev.Name(), lIP.String(), hutils.IPV4MaskToString(lNet.Mask))
	if err != nil {
		return nil, err
	}

	return &WinTunDevice{
		ifc:        tunDev,
		ip:         lIP,
		ipNet:      lNet,
		realMacMap: make(map[string]net.HardwareAddr),
	}, nil
}

func TunClientExtInit(isTargetVPN bool, logger pkg.Logger) {
	logger.Warnf("no op for: %v", isTargetVPN)
}
