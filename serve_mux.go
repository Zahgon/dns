package dns

import (
	"sync"
)

type ServeMux struct {
	z map[string]Handler
	m sync.RWMutex
}

func NewServeMux() *ServeMux { _ = "STUB: not implemented"; return nil }

var DefaultServeMux = NewServeMux()

func (mux *ServeMux) match(q string, t uint16) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

func (mux *ServeMux) Handle(pattern string, handler Handler) { _ = "STUB: not implemented"; return }

func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Msg)) {
	_ = "STUB: not implemented"
	return
}

func (mux *ServeMux) HandleRemove(pattern string) { _ = "STUB: not implemented"; return }

func (mux *ServeMux) ServeDNS(w ResponseWriter, req *Msg) { _ = "STUB: not implemented"; return }

func Handle(pattern string, handler Handler) { _ = "STUB: not implemented"; return }

func HandleRemove(pattern string) { _ = "STUB: not implemented"; return }

func HandleFunc(pattern string, handler func(ResponseWriter, *Msg)) {
	_ = "STUB: not implemented"
	return
}
