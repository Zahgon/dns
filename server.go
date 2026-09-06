package dns

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"sync"
	"time"
)

const maxTCPQueries = 128

var aLongTimeAgo = time.Unix(1, 0)

type Handler interface {
	ServeDNS(w ResponseWriter, r *Msg)
}

type HandlerFunc func(ResponseWriter, *Msg)

func (f HandlerFunc) ServeDNS(w ResponseWriter, r *Msg) { _ = "STUB: not implemented"; return }

type ResponseWriter interface {
	LocalAddr() net.Addr

	RemoteAddr() net.Addr

	WriteMsg(*Msg) error

	Write([]byte) (int, error)

	Close() error

	TsigStatus() error

	TsigTimersOnly(bool)

	Hijack()
}

type ConnectionStater interface {
	ConnectionState() *tls.ConnectionState
}

type response struct {
	closed         bool
	hijacked       bool
	tsigTimersOnly bool
	tsigStatus     error
	tsigRequestMAC string
	tsigProvider   TsigProvider
	udp            net.PacketConn
	tcp            net.Conn
	udpSession     *SessionUDP
	pcSession      net.Addr
	writer         Writer
}

func handleRefused(w ResponseWriter, r *Msg) { _ = "STUB: not implemented"; return }

func HandleFailed(w ResponseWriter, r *Msg) { _ = "STUB: not implemented"; return }

func ListenAndServe(addr string, network string, handler Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func ActivateAndServe(l net.Listener, p net.PacketConn, handler Handler) error {
	_ = "STUB: not implemented"
	return nil
}

type Writer interface {
	io.Writer
}

type Reader interface {
	ReadTCP(conn net.Conn, timeout time.Duration) ([]byte, error)

	ReadUDP(conn *net.UDPConn, timeout time.Duration) ([]byte, *SessionUDP, error)
}

type PacketConnReader interface {
	Reader

	ReadPacketConn(conn net.PacketConn, timeout time.Duration) ([]byte, net.Addr, error)
}

type defaultReader struct {
	*Server
}

var _ PacketConnReader = defaultReader{}

func (dr defaultReader) ReadTCP(conn net.Conn, timeout time.Duration) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dr defaultReader) ReadUDP(conn *net.UDPConn, timeout time.Duration) ([]byte, *SessionUDP, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (dr defaultReader) ReadPacketConn(conn net.PacketConn, timeout time.Duration) ([]byte, net.Addr, error) {
	_ = "STUB: not implemented"
	return nil, *new(net.Addr), nil
}

type DecorateReader func(Reader) Reader

type DecorateWriter func(Writer) Writer

type MsgInvalidFunc func(m []byte, err error)

var DefaultMsgInvalidFunc MsgInvalidFunc = defaultMsgInvalidFunc

func defaultMsgInvalidFunc(m []byte, err error) { _ = "STUB: not implemented"; return }

type Server struct {
	Addr string

	Net string

	Listener net.Listener

	TLSConfig *tls.Config

	PacketConn net.PacketConn

	Handler Handler

	UDPSize int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout func() time.Duration

	TsigProvider TsigProvider

	TsigSecret map[string]string

	NotifyStartedFunc func()

	DecorateReader DecorateReader

	DecorateWriter DecorateWriter

	MaxTCPQueries int

	ReusePort bool

	ReuseAddr bool

	MsgAcceptFunc MsgAcceptFunc

	MsgInvalidFunc MsgInvalidFunc

	lock     sync.RWMutex
	started  bool
	shutdown chan struct{}
	conns    map[net.Conn]struct{}

	udpPool sync.Pool
}

func (srv *Server) tsigProvider() TsigProvider {
	_ = "STUB: not implemented"
	return *new(TsigProvider)
}

func (srv *Server) isStarted() bool { _ = "STUB: not implemented"; return false }

func makeUDPBuffer(size int) func() interface{} { _ = "STUB: not implemented"; return nil }

func (srv *Server) init() {
	srv.shutdown = make(chan struct{})
	srv.conns = make(map[net.Conn]struct{})

	if srv.UDPSize == 0 {
		srv.UDPSize = MinMsgSize
	}
	if srv.MsgAcceptFunc == nil {
		srv.MsgAcceptFunc = DefaultMsgAcceptFunc
	}
	if srv.MsgInvalidFunc == nil {
		srv.MsgInvalidFunc = DefaultMsgInvalidFunc
	}
	if srv.Handler == nil {
		srv.Handler = DefaultServeMux
	}

	srv.udpPool.New = makeUDPBuffer(srv.UDPSize)
}

func unlockOnce(l sync.Locker) func() { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenAndServe() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) ActivateAndServe() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (srv *Server) ShutdownContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

var testShutdownNotify *sync.Cond

func (srv *Server) getReadTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (srv *Server) serveTCP(l net.Listener) error { _ = "STUB: not implemented"; return nil }

func (srv *Server) serveUDP(l net.PacketConn) error { _ = "STUB: not implemented"; return nil }

func (srv *Server) serveTCPConn(wg *sync.WaitGroup, rw net.Conn) { _ = "STUB: not implemented"; return }

func (srv *Server) serveUDPPacket(wg *sync.WaitGroup, m []byte, u net.PacketConn, udpSession *SessionUDP, pcSession net.Addr) {
	_ = "STUB: not implemented"
	return
}

func (srv *Server) serveDNS(m []byte, w *response) { _ = "STUB: not implemented"; return }

func (srv *Server) readTCP(conn net.Conn, timeout time.Duration) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (srv *Server) readUDP(conn *net.UDPConn, timeout time.Duration) ([]byte, *SessionUDP, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (srv *Server) readPacketConn(conn net.PacketConn, timeout time.Duration) ([]byte, net.Addr, error) {
	_ = "STUB: not implemented"
	return nil, *new(net.Addr), nil
}

func (w *response) WriteMsg(m *Msg) (err error) { _ = "STUB: not implemented"; return nil }

func (w *response) Write(m []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *response) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (w *response) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (w *response) TsigStatus() error { _ = "STUB: not implemented"; return nil }

func (w *response) TsigTimersOnly(b bool) { _ = "STUB: not implemented"; return }

func (w *response) Hijack() { _ = "STUB: not implemented"; return }

func (w *response) Close() error { _ = "STUB: not implemented"; return nil }

func (w *response) ConnectionState() *tls.ConnectionState { _ = "STUB: not implemented"; return nil }
