package tun

import "io"

type TunDevice interface {
	io.Reader
	io.Writer
	io.Closer
	Name() string
	RouteAdd(cidr string) error
}
