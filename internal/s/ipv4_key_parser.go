package s

import (
	"errors"
	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
	udpchannel "github.com/jiuzhou-zhao/udp-channel"
	"strings"
)

type IPV4KeyParser struct{}

func NewIPV4KeyParser() *IPV4KeyParser {
	return &IPV4KeyParser{}
}

func (parser *IPV4KeyParser) ParseData(d []byte) (key string, dd []byte, err error) {
	ipPackage := hutils.IPPacket(d)
	if ipPackage.IPver() != 4 {
		err = errors.New("not ip v4")
		return
	}
	key = ipPackage.DstV4().String()
	dd = d
	return
}

func (parser *IPV4KeyParser) ParseKeyFromIPOrCIDR(s string) (key string, err error) {
	cidr, err := udpchannel.ToCIDR(s)
	if err != nil {
		return
	}

	key = strings.Split(cidr, "/")[0]

	return
}

func (parser *IPV4KeyParser) CompareKeyWithCIDR(key string, cidr string) bool {
	if cidr == "" {
		return false
	}
	cidrKey, _ := parser.ParseKeyFromIPOrCIDR(cidr)

	return key == cidrKey
}
