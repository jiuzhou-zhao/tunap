package tun

import "io"

type Device interface {
	io.Reader
	io.Writer
	io.Closer
	Name() string
	RouteAdd(cidr string) error
	RouteDel(cidr string) error
}
