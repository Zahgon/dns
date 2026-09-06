//go:build !windows && !darwin
// +build !windows,!darwin

package dns

import (
	"net"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

var udpOOBSize = func() int {

	oob4 := ipv4.NewControlMessage(ipv4.FlagDst | ipv4.FlagInterface)
	oob6 := ipv6.NewControlMessage(ipv6.FlagDst | ipv6.FlagInterface)

	if len(oob4) > len(oob6) {
		return len(oob4)
	}

	return len(oob6)
}()

type SessionUDP struct {
	raddr   *net.UDPAddr
	context []byte
}

func (s *SessionUDP) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func ReadFromSessionUDP(conn *net.UDPConn, b []byte) (int, *SessionUDP, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func WriteToSessionUDP(conn *net.UDPConn, b []byte, session *SessionUDP) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func setUDPSocketOptions(conn *net.UDPConn) error { _ = "STUB: not implemented"; return nil }

func parseDstFromOOB(oob []byte) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func correctSource(oob []byte) []byte { _ = "STUB: not implemented"; return nil }
