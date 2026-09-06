package dns

const (
	HmacSHA1   = "hmac-sha1."
	HmacSHA224 = "hmac-sha224."
	HmacSHA256 = "hmac-sha256."
	HmacSHA384 = "hmac-sha384."
	HmacSHA512 = "hmac-sha512."

	HmacMD5 = "hmac-md5.sig-alg.reg.int."
)

type TsigProvider interface {
	Generate(msg []byte, t *TSIG) ([]byte, error)

	Verify(msg []byte, t *TSIG) error
}

type tsigHMACProvider string

func (key tsigHMACProvider) Generate(msg []byte, t *TSIG) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (key tsigHMACProvider) Verify(msg []byte, t *TSIG) error {
	_ = "STUB: not implemented"
	return nil
}

type tsigSecretProvider map[string]string

func (ts tsigSecretProvider) Generate(msg []byte, t *TSIG) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts tsigSecretProvider) Verify(msg []byte, t *TSIG) error {
	_ = "STUB: not implemented"
	return nil
}

type TSIG struct {
	Hdr        RR_Header
	Algorithm  string `dns:"domain-name"`
	TimeSigned uint64 `dns:"uint48"`
	Fudge      uint16
	MACSize    uint16
	MAC        string `dns:"size-hex:MACSize"`
	OrigId     uint16
	Error      uint16
	OtherLen   uint16
	OtherData  string `dns:"size-hex:OtherLen"`
}

func (rr *TSIG) String() string { _ = "STUB: not implemented"; return "" }

func (*TSIG) parse(c *zlexer, origin string) *ParseError { _ = "STUB: not implemented"; return nil }

type tsigWireFmt struct {
	Name  string `dns:"domain-name"`
	Class uint16
	Ttl   uint32

	Algorithm  string `dns:"domain-name"`
	TimeSigned uint64 `dns:"uint48"`
	Fudge      uint16

	Error     uint16
	OtherLen  uint16
	OtherData string `dns:"size-hex:OtherLen"`
}

type macWireFmt struct {
	MACSize uint16
	MAC     string `dns:"size-hex:MACSize"`
}

type timerWireFmt struct {
	TimeSigned uint64 `dns:"uint48"`
	Fudge      uint16
}

func TsigGenerate(m *Msg, secret, requestMAC string, timersOnly bool) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func TsigGenerateWithProvider(m *Msg, provider TsigProvider, requestMAC string, timersOnly bool) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func TsigVerify(msg []byte, secret, requestMAC string, timersOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func TsigVerifyWithProvider(msg []byte, provider TsigProvider, requestMAC string, timersOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func tsigVerify(msg []byte, provider TsigProvider, requestMAC string, timersOnly bool, now uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func tsigBuffer(msgbuf []byte, rr *TSIG, requestMAC string, timersOnly bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stripTsig(msg []byte) ([]byte, *TSIG, error) { _ = "STUB: not implemented"; return nil, nil, nil }

func tsigTimeToString(t uint64) string { _ = "STUB: not implemented"; return "" }

func packTsigWire(tw *tsigWireFmt, msg []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packMacWire(mw *macWireFmt, msg []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func packTimerWire(tw *timerWireFmt, msg []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
