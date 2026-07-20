package dns

//go:generate go run msg_generate.go

import (
	"fmt"
	"math/big"
)

const (
	maxCompressionOffset    = 2 << 13
	maxDomainNameWireOctets = 255

	maxCompressionPointers = (maxDomainNameWireOctets+1)/2 - 2

	maxDomainNamePresentationLength = 61*4 + 1 + 63*4 + 1 + 63*4 + 1 + 63*4 + 1
)

var (
	ErrAlg           error = &Error{err: "bad algorithm"}
	ErrAuth          error = &Error{err: "bad authentication"}
	ErrBuf           error = &Error{err: "buffer size too small"}
	ErrConnEmpty     error = &Error{err: "conn has no connection"}
	ErrExtendedRcode error = &Error{err: "bad extended rcode"}
	ErrFqdn          error = &Error{err: "domain must be fully qualified"}
	ErrId            error = &Error{err: "id mismatch"}
	ErrKeyAlg        error = &Error{err: "bad key algorithm"}
	ErrKey           error = &Error{err: "bad key"}
	ErrKeySize       error = &Error{err: "bad key size"}
	ErrLongDomain    error = &Error{err: fmt.Sprintf("domain name exceeded %d wire-format octets", maxDomainNameWireOctets)}
	ErrNoSig         error = &Error{err: "no signature found"}
	ErrPrivKey       error = &Error{err: "bad private key"}
	ErrRcode         error = &Error{err: "bad rcode"}
	ErrRdata         error = &Error{err: "bad rdata"}
	ErrRRset         error = &Error{err: "bad rrset"}
	ErrSecret        error = &Error{err: "no secrets defined"}
	ErrShortRead     error = &Error{err: "short read"}
	ErrSig           error = &Error{err: "bad signature"}
	ErrSoa           error = &Error{err: "no SOA"}
	ErrTime          error = &Error{err: "bad time"}
)

var Id = id

func id() uint16 { _ = "STUB: not implemented"; return 0 }

type MsgHdr struct {
	Id                 uint16
	Response           bool
	Opcode             int
	Authoritative      bool
	Truncated          bool
	RecursionDesired   bool
	RecursionAvailable bool
	Zero               bool
	AuthenticatedData  bool
	CheckingDisabled   bool
	Rcode              int
}

type Msg struct {
	MsgHdr
	Compress bool `json:"-"`
	Question []Question
	Answer   []RR
	Ns       []RR
	Extra    []RR
}

var ClassToString = map[uint16]string{
	ClassINET:   "IN",
	ClassCSNET:  "CS",
	ClassCHAOS:  "CH",
	ClassHESIOD: "HS",
	ClassNONE:   "NONE",
	ClassANY:    "ANY",
}

var OpcodeToString = map[int]string{
	OpcodeQuery:  "QUERY",
	OpcodeIQuery: "IQUERY",
	OpcodeStatus: "STATUS",
	OpcodeNotify: "NOTIFY",
	OpcodeUpdate: "UPDATE",
}

var RcodeToString = map[int]string{
	RcodeSuccess:                    "NOERROR",
	RcodeFormatError:                "FORMERR",
	RcodeServerFailure:              "SERVFAIL",
	RcodeNameError:                  "NXDOMAIN",
	RcodeNotImplemented:             "NOTIMP",
	RcodeRefused:                    "REFUSED",
	RcodeYXDomain:                   "YXDOMAIN",
	RcodeYXRrset:                    "YXRRSET",
	RcodeNXRrset:                    "NXRRSET",
	RcodeNotAuth:                    "NOTAUTH",
	RcodeNotZone:                    "NOTZONE",
	RcodeStatefulTypeNotImplemented: "DSOTYPENI",
	RcodeBadSig:                     "BADSIG",

	RcodeBadKey:    "BADKEY",
	RcodeBadTime:   "BADTIME",
	RcodeBadMode:   "BADMODE",
	RcodeBadName:   "BADNAME",
	RcodeBadAlg:    "BADALG",
	RcodeBadTrunc:  "BADTRUNC",
	RcodeBadCookie: "BADCOOKIE",
}

type compressionMap struct {
	ext map[string]int
	int map[string]uint16
}

func (m compressionMap) valid() bool { _ = "STUB: not implemented"; return false }

func (m compressionMap) insert(s string, pos int) { _ = "STUB: not implemented"; return }

func (m compressionMap) find(s string) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func PackDomainName(s string, msg []byte, off int, compression map[string]int, compress bool) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packDomainName(s string, msg []byte, off int, compression compressionMap, compress bool) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func isRootLabel(s string, bs []byte, off, end int) bool { _ = "STUB: not implemented"; return false }

func UnpackDomainName(msg []byte, off int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packTxt(txt []string, msg []byte, offset int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packTxtString(s string, msg []byte, offset int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packOctetString(s string, msg []byte, offset int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackTxt(msg []byte, off0 int) (ss []string, off int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func isDigit(b byte) bool { _ = "STUB: not implemented"; return false }

func isDDD[T ~[]byte | ~string](s T) bool { _ = "STUB: not implemented"; return false }

func dddToByte[T ~[]byte | ~string](s T) byte { _ = "STUB: not implemented"; return 0 }

func intToBytes(i *big.Int, length int) []byte { _ = "STUB: not implemented"; return nil }

func PackRR(rr RR, msg []byte, off int, compression map[string]int, compress bool) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packRR(rr RR, msg []byte, off int, compression compressionMap, compress bool) (headerEnd int, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func UnpackRR(msg []byte, off int) (rr RR, off1 int, err error) {
	_ = "STUB: not implemented"
	return *new(RR), 0, nil
}

func UnpackRRWithHeader(h RR_Header, msg []byte, off int) (rr RR, off1 int, err error) {
	_ = "STUB: not implemented"
	return *new(RR), 0, nil
}

func unpackRRslice(l int, msg []byte, off int) (dst1 []RR, off1 int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (h *MsgHdr) String() string { _ = "STUB: not implemented"; return "" }

func (dns *Msg) Pack() (msg []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (dns *Msg) PackBuffer(buf []byte) (msg []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dns *Msg) packBufferWithCompressionMap(buf []byte, compression compressionMap, compress bool) (msg []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dns *Msg) unpack(dh Header, msg []byte, off int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (dns *Msg) Unpack(msg []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (dns *Msg) String() string { _ = "STUB: not implemented"; return "" }

func (dns *Msg) isCompressible() bool { _ = "STUB: not implemented"; return false }

func (dns *Msg) Len() int { _ = "STUB: not implemented"; return 0 }

func msgLenWithCompressionMap(dns *Msg, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func domainNameLen(s string, off int, compression map[string]struct{}, compress bool) int {
	_ = "STUB: not implemented"
	return 0
}

func escapedNameLen(s string) int { _ = "STUB: not implemented"; return 0 }

func compressionLenSearch(c map[string]struct{}, s string, msgOff int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func Copy(r RR) RR { _ = "STUB: not implemented"; return *new(RR) }

func Len(r RR) int { _ = "STUB: not implemented"; return 0 }

func (dns *Msg) Copy() *Msg { _ = "STUB: not implemented"; return nil }

func (dns *Msg) CopyTo(r1 *Msg) *Msg { _ = "STUB: not implemented"; return nil }

func (q *Question) pack(msg []byte, off int, compression compressionMap, compress bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackQuestion(msg []byte, off int) (Question, int, error) {
	_ = "STUB: not implemented"
	return *new(Question), 0, nil
}

func (dh *Header) pack(msg []byte, off int, compression compressionMap, compress bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackMsgHdr(msg []byte, off int) (Header, int, error) {
	_ = "STUB: not implemented"
	return *new(Header), 0, nil
}

func (dns *Msg) setHdr(dh Header) { _ = "STUB: not implemented"; return }
