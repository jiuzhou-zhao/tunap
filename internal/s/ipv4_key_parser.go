package s

import (
	"errors"
	"github.com/jiuzhou-zhao/tunap/pkg/hutils"
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
