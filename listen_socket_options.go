//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build aix darwin dragonfly freebsd linux netbsd openbsd

package dns

import (
	"net"
	"syscall"
)

const supportsReusePort = true

func reuseportControl(network, address string, c syscall.RawConn) error {
	_ = "STUB: not implemented"
	return nil
}

const supportsReuseAddr = true

func reuseaddrControl(network, address string, c syscall.RawConn) error {
	_ = "STUB: not implemented"
	return nil
}

func reuseaddrandportControl(network, address string, c syscall.RawConn) error {
	_ = "STUB: not implemented"
	return nil
}

func checkReuseport(fd uintptr) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func checkReuseaddr(fd uintptr) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func listenTCP(network, addr string, reuseport, reuseaddr bool) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func listenUDP(network, addr string, reuseport, reuseaddr bool) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}
