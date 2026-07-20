package dns

import (
	"context"
	"crypto/tls"
	"net"
	"time"
)

const (
	dnsTimeout     time.Duration = 2 * time.Second
	tcpIdleTimeout time.Duration = 8 * time.Second
)

func isPacketConn(c net.Conn) bool { _ = "STUB: not implemented"; return false }

type Conn struct {
	net.Conn
	UDPSize        uint16
	TsigSecret     map[string]string
	TsigProvider   TsigProvider
	tsigRequestMAC string
}

func (co *Conn) tsigProvider() TsigProvider { _ = "STUB: not implemented"; return *new(TsigProvider) }

type Client struct {
	Net       string
	UDPSize   uint16
	TLSConfig *tls.Config
	Dialer    *net.Dialer

	Timeout      time.Duration
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TsigSecret   map[string]string
	TsigProvider TsigProvider

	SingleInflight bool
}

func Exchange(m *Msg, a string) (r *Msg, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) dialTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *Client) readTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (c *Client) writeTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Client) Dial(address string) (conn *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) DialContext(ctx context.Context, address string) (conn *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Exchange(m *Msg, address string) (r *Msg, rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration), nil
}

func (c *Client) ExchangeWithConn(m *Msg, conn *Conn) (r *Msg, rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration), nil
}

func (c *Client) ExchangeWithConnContext(ctx context.Context, m *Msg, co *Conn) (r *Msg, rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration), nil
}

func (co *Conn) ReadMsg() (*Msg, error) { _ = "STUB: not implemented"; return nil, nil }

func (co *Conn) ReadMsgHeader(hdr *Header) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (co *Conn) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (co *Conn) WriteMsg(m *Msg) (err error) { _ = "STUB: not implemented"; return nil }

func (co *Conn) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Client) getTimeoutForRequest(timeout time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func Dial(network, address string) (conn *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExchangeContext(ctx context.Context, m *Msg, a string) (r *Msg, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExchangeConn(c net.Conn, m *Msg) (r *Msg, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialTimeout(network, address string, timeout time.Duration) (conn *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialWithTLS(network, address string, tlsConfig *tls.Config) (conn *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialTimeoutWithTLS(network, address string, tlsConfig *tls.Config, timeout time.Duration) (conn *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ExchangeContext(ctx context.Context, m *Msg, a string) (r *Msg, rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration), nil
}
