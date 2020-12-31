package hutils

import (
	"net"
	"strconv"
	"strings"
)

func IPV4MaskToString(ipMask net.IPMask) string {
	var s []string
	for _, i := range ipMask[:] {
		s = append(s, strconv.Itoa(int(i)))
	}
	return strings.Join(s, ".")
}
