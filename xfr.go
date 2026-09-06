package dns

import (
	"crypto/tls"
	"time"
)

type Envelope struct {
	RR    []RR
	Error error
}

type Transfer struct {
	*Conn
	DialTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	TsigProvider   TsigProvider
	TsigSecret     map[string]string
	tsigTimersOnly bool
	TLS            *tls.Config
}

func (t *Transfer) tsigProvider() TsigProvider {
	_ = "STUB: not implemented"
	return *new(TsigProvider)
}

func (t *Transfer) In(q *Msg, a string) (env chan *Envelope, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transfer) inAxfr(q *Msg, c chan *Envelope) { _ = "STUB: not implemented"; return }

func (t *Transfer) inIxfr(q *Msg, c chan *Envelope) { _ = "STUB: not implemented"; return }

func (t *Transfer) Out(w ResponseWriter, q *Msg, ch chan *Envelope) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transfer) ReadMsg() (*Msg, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Transfer) WriteMsg(m *Msg) (err error) { _ = "STUB: not implemented"; return nil }

func isSOAFirst(in *Msg) bool { _ = "STUB: not implemented"; return false }

func isSOALast(in *Msg) bool { _ = "STUB: not implemented"; return false }

const errXFR = "bad xfr rcode: %d"
