package hutils

import (
	"github.com/mdlayher/arp"
	"github.com/mdlayher/ethernet"
)

func ParseEthernetFrame(buf []byte) (*arp.Packet, *ethernet.Frame, error) {
	f := new(ethernet.Frame)
	if err := f.UnmarshalBinary(buf); err != nil {
		return nil, nil, err
	}

	var p *arp.Packet

	if f.EtherType == ethernet.EtherTypeARP {
		p = new(arp.Packet)
		if err := p.UnmarshalBinary(f.Payload); err != nil {
			return nil, nil, err
		}
	}

	return p, f, nil
}
