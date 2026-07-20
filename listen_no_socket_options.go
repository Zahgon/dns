//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd

package dns

import (
	"net"
)

const (
	supportsReusePort = false
	supportsReuseAddr = false
)

func listenTCP(network, addr string, reuseport, reuseaddr bool) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func listenUDP(network, addr string, reuseport, reuseaddr bool) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

func checkReuseport(fd uintptr) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func checkReuseaddr(fd uintptr) (bool, error) { _ = "STUB: not implemented"; return false, nil }
