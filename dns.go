package dns

const (
	year68     = 1 << 31
	defaultTtl = 3600

	DefaultMsgSize = 4096

	MinMsgSize = 512

	MaxMsgSize = 65535
)

type Error struct{ err string }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

type RR interface {
	Header() *RR_Header

	String() string

	copy() RR

	len(off int, compression map[string]struct{}) int

	pack(msg []byte, off int, compression compressionMap, compress bool) (off1 int, err error)

	unpack(msg []byte, off int) (off1 int, err error)

	parse(c *zlexer, origin string) *ParseError

	isDuplicate(r2 RR) bool
}

type RR_Header struct {
	Name     string `dns:"cdomain-name"`
	Rrtype   uint16
	Class    uint16
	Ttl      uint32
	Rdlength uint16
}

func (h *RR_Header) Header() *RR_Header { _ = "STUB: not implemented"; return nil }

func (h *RR_Header) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (h *RR_Header) String() string { _ = "STUB: not implemented"; return "" }

func (h *RR_Header) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (h *RR_Header) pack(msg []byte, off int, compression compressionMap, compress bool) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *RR_Header) unpack(msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *RR_Header) parse(c *zlexer, origin string) *ParseError {
	_ = "STUB: not implemented"
	return nil
}

func (rr *RFC3597) ToRFC3597(r RR) error { _ = "STUB: not implemented"; return nil }

func (rr *RFC3597) fromRFC3597(r RR) error { _ = "STUB: not implemented"; return nil }
