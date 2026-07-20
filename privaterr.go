package dns

type PrivateRdata interface {
	String() string

	Parse([]string) error

	Pack([]byte) (int, error)

	Unpack([]byte) (int, error)

	Copy(PrivateRdata) error

	Len() int
}

type PrivateRR struct {
	Hdr  RR_Header
	Data PrivateRdata

	generator func() PrivateRdata
}

func (r *PrivateRR) Header() *RR_Header { _ = "STUB: not implemented"; return nil }

func (r *PrivateRR) String() string { _ = "STUB: not implemented"; return "" }

func (r *PrivateRR) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (r *PrivateRR) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (r *PrivateRR) pack(msg []byte, off int, compression compressionMap, compress bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *PrivateRR) unpack(msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *PrivateRR) parse(c *zlexer, origin string) *ParseError {
	_ = "STUB: not implemented"
	return nil
}

func (r *PrivateRR) isDuplicate(r2 RR) bool { _ = "STUB: not implemented"; return false }

func PrivateHandle(rtypestr string, rtype uint16, generator func() PrivateRdata) {
	_ = "STUB: not implemented"
	return
}

func PrivateHandleRemove(rtype uint16) { _ = "STUB: not implemented"; return }
