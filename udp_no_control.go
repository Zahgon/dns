//go:build windows || darwin
// +build windows darwin

package dns

import "net"

type SessionUDP struct {
	raddr *net.UDPAddr
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

func setUDPSocketOptions(*net.UDPConn) error { _ = "STUB: not implemented"; return nil }
func parseDstFromOOB([]byte, net.IP) net.IP  { _ = "STUB: not implemented"; return *new(net.IP) }
