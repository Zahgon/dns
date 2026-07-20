package dns

const hexDigit = "0123456789abcdef"

func (dns *Msg) SetReply(request *Msg) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetQuestion(z string, t uint16) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetNotify(z string) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetRcode(request *Msg, rcode int) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetRcodeFormatError(request *Msg) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetUpdate(z string) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetIxfr(z string, serial uint32, ns, mbox string) *Msg {
	_ = "STUB: not implemented"
	return nil
}

func (dns *Msg) SetAxfr(z string) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) SetTsig(z, algo string, fudge uint16, timesigned int64) *Msg {
	_ = "STUB: not implemented"
	return nil
}

func (dns *Msg) SetEdns0(udpsize uint16, do bool) *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) IsTsig() *TSIG { _ = "STUB: not implemented"; return nil }

func (dns *Msg) IsEdns0() *OPT { _ = "STUB: not implemented"; return nil }

func (dns *Msg) popEdns0() *OPT { _ = "STUB: not implemented"; return nil }

func IsDomainName(s string) (labels int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func IsSubDomain(parent, child string) bool { _ = "STUB: not implemented"; return false }

func IsMsg(buf []byte) error { _ = "STUB: not implemented"; return nil }

func IsFqdn(s string) bool { _ = "STUB: not implemented"; return false }

func IsRRset(rrset []RR) bool { _ = "STUB: not implemented"; return false }

func Fqdn(s string) string { _ = "STUB: not implemented"; return "" }

func CanonicalName(s string) string { _ = "STUB: not implemented"; return "" }

func ReverseAddr(addr string) (arpa string, err error) { _ = "STUB: not implemented"; return "", nil }

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func (c Class) String() string { _ = "STUB: not implemented"; return "" }

func (n Name) String() string { _ = "STUB: not implemented"; return "" }
