package dns

import (
	"net"
)

const (
	EDNS0LLQ          = 0x1
	EDNS0UL           = 0x2
	EDNS0NSID         = 0x3
	EDNS0ESU          = 0x4
	EDNS0DAU          = 0x5
	EDNS0DHU          = 0x6
	EDNS0N3U          = 0x7
	EDNS0SUBNET       = 0x8
	EDNS0EXPIRE       = 0x9
	EDNS0COOKIE       = 0xa
	EDNS0TCPKEEPALIVE = 0xb
	EDNS0PADDING      = 0xc
	EDNS0EDE          = 0xf
	EDNS0REPORTING    = 0x12
	EDNS0ZONEVERSION  = 0x13
	EDNS0LOCALSTART   = 0xFDE9
	EDNS0LOCALEND     = 0xFFFE
	_DO               = 1 << 15
	_CO               = 1 << 14
)

func makeDataOpt(code uint16) EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type OPT struct {
	Hdr    RR_Header
	Option []EDNS0 `dns:"opt"`
}

func (rr *OPT) String() string { _ = "STUB: not implemented"; return "" }

func (rr *OPT) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (*OPT) parse(c *zlexer, origin string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *OPT) isDuplicate(r2 RR) bool { _ = "STUB: not implemented"; return false }

func (rr *OPT) Version() uint8 { _ = "STUB: not implemented"; return 0 }

func (rr *OPT) SetVersion(v uint8) { _ = "STUB: not implemented"; return }

func (rr *OPT) ExtendedRcode() int { _ = "STUB: not implemented"; return 0 }

func (rr *OPT) SetExtendedRcode(v uint16) { _ = "STUB: not implemented"; return }

func (rr *OPT) UDPSize() uint16 { _ = "STUB: not implemented"; return 0 }

func (rr *OPT) SetUDPSize(size uint16) { _ = "STUB: not implemented"; return }

func (rr *OPT) Do() bool { _ = "STUB: not implemented"; return false }

func (rr *OPT) SetDo(do ...bool) { _ = "STUB: not implemented"; return }

func (rr *OPT) Co() bool { _ = "STUB: not implemented"; return false }

func (rr *OPT) SetCo(co ...bool) { _ = "STUB: not implemented"; return }

func (rr *OPT) Z() uint16 { _ = "STUB: not implemented"; return 0 }

func (rr *OPT) SetZ(z uint16) { _ = "STUB: not implemented"; return }

type EDNS0 interface {
	Option() uint16

	pack() ([]byte, error)

	unpack([]byte) error

	String() string

	copy() EDNS0
}

type EDNS0_NSID struct {
	Code uint16
	Nsid string
}

func (e *EDNS0_NSID) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_NSID) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_NSID) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }
func (e *EDNS0_NSID) String() string        { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_NSID) copy() EDNS0           { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_SUBNET struct {
	Code          uint16
	Family        uint16
	SourceNetmask uint8
	SourceScope   uint8
	Address       net.IP
}

func (e *EDNS0_SUBNET) Option() uint16 { _ = "STUB: not implemented"; return 0 }

func (e *EDNS0_SUBNET) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_SUBNET) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_SUBNET) String() (s string) { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_SUBNET) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_COOKIE struct {
	Code   uint16
	Cookie string
}

func (e *EDNS0_COOKIE) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_COOKIE) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_COOKIE) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }
func (e *EDNS0_COOKIE) String() string        { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_COOKIE) copy() EDNS0           { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_UL struct {
	Code     uint16
	Lease    uint32
	KeyLease uint32
}

func (e *EDNS0_UL) Option() uint16 { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_UL) String() string { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_UL) copy() EDNS0    { _ = "STUB: not implemented"; return *new(EDNS0) }

func (e *EDNS0_UL) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_UL) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

type EDNS0_LLQ struct {
	Code      uint16
	Version   uint16
	Opcode    uint16
	Error     uint16
	Id        uint64
	LeaseLife uint32
}

func (e *EDNS0_LLQ) Option() uint16 { _ = "STUB: not implemented"; return 0 }

func (e *EDNS0_LLQ) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_LLQ) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_LLQ) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_LLQ) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_DAU struct {
	Code    uint16
	AlgCode []uint8
}

func (e *EDNS0_DAU) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_DAU) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (e *EDNS0_DAU) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_DAU) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_DAU) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_DHU struct {
	Code    uint16
	AlgCode []uint8
}

func (e *EDNS0_DHU) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_DHU) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (e *EDNS0_DHU) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_DHU) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_DHU) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_N3U struct {
	Code    uint16
	AlgCode []uint8
}

func (e *EDNS0_N3U) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_N3U) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (e *EDNS0_N3U) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_N3U) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_N3U) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_EXPIRE struct {
	Code   uint16
	Expire uint32
	Empty  bool
}

func (e *EDNS0_EXPIRE) Option() uint16 { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_EXPIRE) copy() EDNS0    { _ = "STUB: not implemented"; return *new(EDNS0) }

func (e *EDNS0_EXPIRE) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_EXPIRE) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_EXPIRE) String() (s string) { _ = "STUB: not implemented"; return "" }

type EDNS0_LOCAL struct {
	Code uint16
	Data []byte
}

func (e *EDNS0_LOCAL) Option() uint16 { _ = "STUB: not implemented"; return 0 }

func (e *EDNS0_LOCAL) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_LOCAL) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

func (e *EDNS0_LOCAL) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_LOCAL) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

type EDNS0_TCP_KEEPALIVE struct {
	Code uint16

	Timeout uint16

	Length uint16
}

func (e *EDNS0_TCP_KEEPALIVE) Option() uint16 { _ = "STUB: not implemented"; return 0 }

func (e *EDNS0_TCP_KEEPALIVE) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_TCP_KEEPALIVE) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *EDNS0_TCP_KEEPALIVE) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_TCP_KEEPALIVE) copy() EDNS0 { _ = "STUB: not implemented"; return *new(EDNS0) }

type EDNS0_PADDING struct {
	Padding []byte
}

func (e *EDNS0_PADDING) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_PADDING) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (e *EDNS0_PADDING) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }
func (e *EDNS0_PADDING) String() string        { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_PADDING) copy() EDNS0           { _ = "STUB: not implemented"; return *new(EDNS0) }

const (
	ExtendedErrorCodeOther uint16 = iota
	ExtendedErrorCodeUnsupportedDNSKEYAlgorithm
	ExtendedErrorCodeUnsupportedDSDigestType
	ExtendedErrorCodeStaleAnswer
	ExtendedErrorCodeForgedAnswer
	ExtendedErrorCodeDNSSECIndeterminate
	ExtendedErrorCodeDNSBogus
	ExtendedErrorCodeSignatureExpired
	ExtendedErrorCodeSignatureNotYetValid
	ExtendedErrorCodeDNSKEYMissing
	ExtendedErrorCodeRRSIGsMissing
	ExtendedErrorCodeNoZoneKeyBitSet
	ExtendedErrorCodeNSECMissing
	ExtendedErrorCodeCachedError
	ExtendedErrorCodeNotReady
	ExtendedErrorCodeBlocked
	ExtendedErrorCodeCensored
	ExtendedErrorCodeFiltered
	ExtendedErrorCodeProhibited
	ExtendedErrorCodeStaleNXDOMAINAnswer
	ExtendedErrorCodeNotAuthoritative
	ExtendedErrorCodeNotSupported
	ExtendedErrorCodeNoReachableAuthority
	ExtendedErrorCodeNetworkError
	ExtendedErrorCodeInvalidData
	ExtendedErrorCodeSignatureExpiredBeforeValid
	ExtendedErrorCodeTooEarly
	ExtendedErrorCodeUnsupportedNSEC3IterValue
	ExtendedErrorCodeUnableToConformToPolicy
	ExtendedErrorCodeSynthesized
	ExtendedErrorCodeInvalidQueryType
)

var ExtendedErrorCodeToString = map[uint16]string{
	ExtendedErrorCodeOther:                       "Other",
	ExtendedErrorCodeUnsupportedDNSKEYAlgorithm:  "Unsupported DNSKEY Algorithm",
	ExtendedErrorCodeUnsupportedDSDigestType:     "Unsupported DS Digest Type",
	ExtendedErrorCodeStaleAnswer:                 "Stale Answer",
	ExtendedErrorCodeForgedAnswer:                "Forged Answer",
	ExtendedErrorCodeDNSSECIndeterminate:         "DNSSEC Indeterminate",
	ExtendedErrorCodeDNSBogus:                    "DNSSEC Bogus",
	ExtendedErrorCodeSignatureExpired:            "Signature Expired",
	ExtendedErrorCodeSignatureNotYetValid:        "Signature Not Yet Valid",
	ExtendedErrorCodeDNSKEYMissing:               "DNSKEY Missing",
	ExtendedErrorCodeRRSIGsMissing:               "RRSIGs Missing",
	ExtendedErrorCodeNoZoneKeyBitSet:             "No Zone Key Bit Set",
	ExtendedErrorCodeNSECMissing:                 "NSEC Missing",
	ExtendedErrorCodeCachedError:                 "Cached Error",
	ExtendedErrorCodeNotReady:                    "Not Ready",
	ExtendedErrorCodeBlocked:                     "Blocked",
	ExtendedErrorCodeCensored:                    "Censored",
	ExtendedErrorCodeFiltered:                    "Filtered",
	ExtendedErrorCodeProhibited:                  "Prohibited",
	ExtendedErrorCodeStaleNXDOMAINAnswer:         "Stale NXDOMAIN Answer",
	ExtendedErrorCodeNotAuthoritative:            "Not Authoritative",
	ExtendedErrorCodeNotSupported:                "Not Supported",
	ExtendedErrorCodeNoReachableAuthority:        "No Reachable Authority",
	ExtendedErrorCodeNetworkError:                "Network Error",
	ExtendedErrorCodeInvalidData:                 "Invalid Data",
	ExtendedErrorCodeSignatureExpiredBeforeValid: "Signature Expired Before Valid",
	ExtendedErrorCodeTooEarly:                    "Too Early",
	ExtendedErrorCodeUnsupportedNSEC3IterValue:   "Unsupported NSEC3 Iterations Value",
	ExtendedErrorCodeUnableToConformToPolicy:     "Unable To Conform To Policy",
	ExtendedErrorCodeSynthesized:                 "Synthesized",
	ExtendedErrorCodeInvalidQueryType:            "Invalid Query Type",
}

var StringToExtendedErrorCode = reverseInt16(ExtendedErrorCodeToString)

type EDNS0_EDE struct {
	InfoCode  uint16
	ExtraText string
}

func (e *EDNS0_EDE) Option() uint16 { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_EDE) copy() EDNS0    { _ = "STUB: not implemented"; return *new(EDNS0) }

func (e *EDNS0_EDE) String() string { _ = "STUB: not implemented"; return "" }

func (e *EDNS0_EDE) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_EDE) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

type EDNS0_ESU struct {
	Code uint16
	Uri  string
}

func (e *EDNS0_ESU) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_ESU) String() string        { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_ESU) copy() EDNS0           { _ = "STUB: not implemented"; return *new(EDNS0) }
func (e *EDNS0_ESU) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (e *EDNS0_ESU) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

type EDNS0_REPORTING struct {
	Code        uint16
	AgentDomain string
}

func (e *EDNS0_REPORTING) Option() uint16        { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_REPORTING) String() string        { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_REPORTING) copy() EDNS0           { _ = "STUB: not implemented"; return *new(EDNS0) }
func (e *EDNS0_REPORTING) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_REPORTING) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

type EDNS0_ZONEVERSION struct {
	Code uint16

	LabelCount uint8

	Type uint8

	Version string
}

func (e *EDNS0_ZONEVERSION) Option() uint16 { _ = "STUB: not implemented"; return 0 }
func (e *EDNS0_ZONEVERSION) String() string { _ = "STUB: not implemented"; return "" }
func (e *EDNS0_ZONEVERSION) copy() EDNS0    { _ = "STUB: not implemented"; return *new(EDNS0) }

func (e *EDNS0_ZONEVERSION) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EDNS0_ZONEVERSION) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }
