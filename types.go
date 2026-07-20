package dns

import (
	"net"
	"strings"
)

type (
	Type uint16

	Class uint16

	Name string
)

const (
	TypeNone       uint16 = 0
	TypeA          uint16 = 1
	TypeNS         uint16 = 2
	TypeMD         uint16 = 3
	TypeMF         uint16 = 4
	TypeCNAME      uint16 = 5
	TypeSOA        uint16 = 6
	TypeMB         uint16 = 7
	TypeMG         uint16 = 8
	TypeMR         uint16 = 9
	TypeNULL       uint16 = 10
	TypePTR        uint16 = 12
	TypeHINFO      uint16 = 13
	TypeMINFO      uint16 = 14
	TypeMX         uint16 = 15
	TypeTXT        uint16 = 16
	TypeRP         uint16 = 17
	TypeAFSDB      uint16 = 18
	TypeX25        uint16 = 19
	TypeISDN       uint16 = 20
	TypeRT         uint16 = 21
	TypeNSAPPTR    uint16 = 23
	TypeSIG        uint16 = 24
	TypeKEY        uint16 = 25
	TypePX         uint16 = 26
	TypeGPOS       uint16 = 27
	TypeAAAA       uint16 = 28
	TypeLOC        uint16 = 29
	TypeNXT        uint16 = 30
	TypeEID        uint16 = 31
	TypeNIMLOC     uint16 = 32
	TypeSRV        uint16 = 33
	TypeATMA       uint16 = 34
	TypeNAPTR      uint16 = 35
	TypeKX         uint16 = 36
	TypeCERT       uint16 = 37
	TypeDNAME      uint16 = 39
	TypeOPT        uint16 = 41
	TypeAPL        uint16 = 42
	TypeDS         uint16 = 43
	TypeSSHFP      uint16 = 44
	TypeIPSECKEY   uint16 = 45
	TypeRRSIG      uint16 = 46
	TypeNSEC       uint16 = 47
	TypeDNSKEY     uint16 = 48
	TypeDHCID      uint16 = 49
	TypeNSEC3      uint16 = 50
	TypeNSEC3PARAM uint16 = 51
	TypeTLSA       uint16 = 52
	TypeSMIMEA     uint16 = 53
	TypeHIP        uint16 = 55
	TypeNINFO      uint16 = 56
	TypeRKEY       uint16 = 57
	TypeTALINK     uint16 = 58
	TypeCDS        uint16 = 59
	TypeCDNSKEY    uint16 = 60
	TypeOPENPGPKEY uint16 = 61
	TypeCSYNC      uint16 = 62
	TypeZONEMD     uint16 = 63
	TypeSVCB       uint16 = 64
	TypeHTTPS      uint16 = 65
	TypeSPF        uint16 = 99
	TypeUINFO      uint16 = 100
	TypeUID        uint16 = 101
	TypeGID        uint16 = 102
	TypeUNSPEC     uint16 = 103
	TypeNID        uint16 = 104
	TypeL32        uint16 = 105
	TypeL64        uint16 = 106
	TypeLP         uint16 = 107
	TypeEUI48      uint16 = 108
	TypeEUI64      uint16 = 109
	TypeNXNAME     uint16 = 128
	TypeURI        uint16 = 256
	TypeCAA        uint16 = 257
	TypeAVC        uint16 = 258
	TypeAMTRELAY   uint16 = 260
	TypeRESINFO    uint16 = 261

	TypeTKEY uint16 = 249
	TypeTSIG uint16 = 250

	TypeIXFR  uint16 = 251
	TypeAXFR  uint16 = 252
	TypeMAILB uint16 = 253
	TypeMAILA uint16 = 254
	TypeANY   uint16 = 255

	TypeTA       uint16 = 32768
	TypeDLV      uint16 = 32769
	TypeReserved uint16 = 65535

	ClassINET   = 1
	ClassCSNET  = 2
	ClassCHAOS  = 3
	ClassHESIOD = 4
	ClassNONE   = 254
	ClassANY    = 255

	RcodeSuccess                    = 0
	RcodeFormatError                = 1
	RcodeServerFailure              = 2
	RcodeNameError                  = 3
	RcodeNotImplemented             = 4
	RcodeRefused                    = 5
	RcodeYXDomain                   = 6
	RcodeYXRrset                    = 7
	RcodeNXRrset                    = 8
	RcodeNotAuth                    = 9
	RcodeNotZone                    = 10
	RcodeStatefulTypeNotImplemented = 11
	RcodeBadSig                     = 16
	RcodeBadVers                    = 16
	RcodeBadKey                     = 17
	RcodeBadTime                    = 18
	RcodeBadMode                    = 19
	RcodeBadName                    = 20
	RcodeBadAlg                     = 21
	RcodeBadTrunc                   = 22
	RcodeBadCookie                  = 23

	OpcodeQuery    = 0
	OpcodeIQuery   = 1
	OpcodeStatus   = 2
	OpcodeNotify   = 4
	OpcodeUpdate   = 5
	OpcodeStateful = 6
)

const (
	ZoneMDSchemeSimple = 1

	ZoneMDHashAlgSHA384 = 1
	ZoneMDHashAlgSHA512 = 2
)

const (
	IPSECGatewayNone uint8 = iota
	IPSECGatewayIPv4
	IPSECGatewayIPv6
	IPSECGatewayHost
)

const (
	AMTRELAYNone = IPSECGatewayNone
	AMTRELAYIPv4 = IPSECGatewayIPv4
	AMTRELAYIPv6 = IPSECGatewayIPv6
	AMTRELAYHost = IPSECGatewayHost
)

const (
	StatefulTypeKeepAlive uint16 = iota + 1
	StatefulTypeRetryDelay
	StatefulTypeEncryptionPadding
)

var StatefulTypeToString = map[uint16]string{
	StatefulTypeKeepAlive:         "KeepAlive",
	StatefulTypeRetryDelay:        "RetryDelay",
	StatefulTypeEncryptionPadding: "EncryptionPadding",
}

type Header struct {
	Id                                 uint16
	Bits                               uint16
	Qdcount, Ancount, Nscount, Arcount uint16
}

const (
	headerSize = 12

	_QR = 1 << 15
	_AA = 1 << 10
	_TC = 1 << 9
	_RD = 1 << 8
	_RA = 1 << 7
	_Z  = 1 << 6
	_AD = 1 << 5
	_CD = 1 << 4
)

const (
	LOC_EQUATOR       = 1 << 31
	LOC_PRIMEMERIDIAN = 1 << 31
	LOC_HOURS         = 60 * 1000
	LOC_DEGREES       = 60 * LOC_HOURS
	LOC_ALTITUDEBASE  = 100000
)

const (
	CertPKIX = 1 + iota
	CertSPKI
	CertPGP
	CertIPIX
	CertISPKI
	CertIPGP
	CertACPKIX
	CertIACPKIX
	CertURI = 253
	CertOID = 254
)

var CertTypeToString = map[uint16]string{
	CertPKIX:    "PKIX",
	CertSPKI:    "SPKI",
	CertPGP:     "PGP",
	CertIPIX:    "IPIX",
	CertISPKI:   "ISPKI",
	CertIPGP:    "IPGP",
	CertACPKIX:  "ACPKIX",
	CertIACPKIX: "IACPKIX",
	CertURI:     "URI",
	CertOID:     "OID",
}

const ipv4InIPv6Prefix = "::ffff:"

//go:generate go run types_generate.go

type Question struct {
	Name   string `dns:"cdomain-name"`
	Qtype  uint16
	Qclass uint16
}

func (q *Question) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (q *Question) String() (s string) { _ = "STUB: not implemented"; return "" }

type ANY struct {
	Hdr RR_Header
}

func (rr *ANY) String() string { _ = "STUB: not implemented"; return "" }

func (*ANY) parse(c *zlexer, origin string) *ParseError { _ = "STUB: not implemented"; return nil }

type NULL struct {
	Hdr  RR_Header
	Data string `dns:"any"`
}

func (rr *NULL) String() string { _ = "STUB: not implemented"; return "" }

func (*NULL) parse(c *zlexer, origin string) *ParseError { _ = "STUB: not implemented"; return nil }

type NXNAME struct {
	Hdr RR_Header
}

func (rr *NXNAME) String() string { _ = "STUB: not implemented"; return "" }

func (*NXNAME) parse(c *zlexer, origin string) *ParseError { _ = "STUB: not implemented"; return nil }

type CNAME struct {
	Hdr    RR_Header
	Target string `dns:"cdomain-name"`
}

func (rr *CNAME) String() string { _ = "STUB: not implemented"; return "" }

type HINFO struct {
	Hdr RR_Header
	Cpu string
	Os  string
}

func (rr *HINFO) String() string { _ = "STUB: not implemented"; return "" }

type MB struct {
	Hdr RR_Header
	Mb  string `dns:"cdomain-name"`
}

func (rr *MB) String() string { _ = "STUB: not implemented"; return "" }

type MG struct {
	Hdr RR_Header
	Mg  string `dns:"cdomain-name"`
}

func (rr *MG) String() string { _ = "STUB: not implemented"; return "" }

type MINFO struct {
	Hdr   RR_Header
	Rmail string `dns:"cdomain-name"`
	Email string `dns:"cdomain-name"`
}

func (rr *MINFO) String() string { _ = "STUB: not implemented"; return "" }

type MR struct {
	Hdr RR_Header
	Mr  string `dns:"cdomain-name"`
}

func (rr *MR) String() string { _ = "STUB: not implemented"; return "" }

type MF struct {
	Hdr RR_Header
	Mf  string `dns:"cdomain-name"`
}

func (rr *MF) String() string { _ = "STUB: not implemented"; return "" }

type MD struct {
	Hdr RR_Header
	Md  string `dns:"cdomain-name"`
}

func (rr *MD) String() string { _ = "STUB: not implemented"; return "" }

type MX struct {
	Hdr        RR_Header
	Preference uint16
	Mx         string `dns:"cdomain-name"`
}

func (rr *MX) String() string { _ = "STUB: not implemented"; return "" }

type AFSDB struct {
	Hdr      RR_Header
	Subtype  uint16
	Hostname string `dns:"domain-name"`
}

func (rr *AFSDB) String() string { _ = "STUB: not implemented"; return "" }

type X25 struct {
	Hdr         RR_Header
	PSDNAddress string
}

func (rr *X25) String() string { _ = "STUB: not implemented"; return "" }

type ISDN struct {
	Hdr        RR_Header
	Address    string
	SubAddress string
}

func (rr *ISDN) String() string { _ = "STUB: not implemented"; return "" }

type RT struct {
	Hdr        RR_Header
	Preference uint16
	Host       string `dns:"domain-name"`
}

func (rr *RT) String() string { _ = "STUB: not implemented"; return "" }

type NS struct {
	Hdr RR_Header
	Ns  string `dns:"cdomain-name"`
}

func (rr *NS) String() string { _ = "STUB: not implemented"; return "" }

type PTR struct {
	Hdr RR_Header
	Ptr string `dns:"cdomain-name"`
}

func (rr *PTR) String() string { _ = "STUB: not implemented"; return "" }

type RP struct {
	Hdr  RR_Header
	Mbox string `dns:"domain-name"`
	Txt  string `dns:"domain-name"`
}

func (rr *RP) String() string { _ = "STUB: not implemented"; return "" }

type SOA struct {
	Hdr     RR_Header
	Ns      string `dns:"cdomain-name"`
	Mbox    string `dns:"cdomain-name"`
	Serial  uint32
	Refresh uint32
	Retry   uint32
	Expire  uint32
	Minttl  uint32
}

func (rr *SOA) String() string { _ = "STUB: not implemented"; return "" }

type TXT struct {
	Hdr RR_Header
	Txt []string `dns:"txt"`
}

func (rr *TXT) String() string { _ = "STUB: not implemented"; return "" }

func sprintName(s string) string { _ = "STUB: not implemented"; return "" }

func sprintTxtOctet(s string) string { _ = "STUB: not implemented"; return "" }

func sprintTxt(txt []string) string { _ = "STUB: not implemented"; return "" }

func writeTXTStringByte(s *strings.Builder, b byte) { _ = "STUB: not implemented"; return }

const (
	escapedByteSmall = "" +
		`\000\001\002\003\004\005\006\007\008\009` +
		`\010\011\012\013\014\015\016\017\018\019` +
		`\020\021\022\023\024\025\026\027\028\029` +
		`\030\031`
	escapedByteLarge = `\127\128\129` +
		`\130\131\132\133\134\135\136\137\138\139` +
		`\140\141\142\143\144\145\146\147\148\149` +
		`\150\151\152\153\154\155\156\157\158\159` +
		`\160\161\162\163\164\165\166\167\168\169` +
		`\170\171\172\173\174\175\176\177\178\179` +
		`\180\181\182\183\184\185\186\187\188\189` +
		`\190\191\192\193\194\195\196\197\198\199` +
		`\200\201\202\203\204\205\206\207\208\209` +
		`\210\211\212\213\214\215\216\217\218\219` +
		`\220\221\222\223\224\225\226\227\228\229` +
		`\230\231\232\233\234\235\236\237\238\239` +
		`\240\241\242\243\244\245\246\247\248\249` +
		`\250\251\252\253\254\255`
)

func escapeByte(b byte) string { _ = "STUB: not implemented"; return "" }

func isDomainNameLabelSpecial(b byte) bool { _ = "STUB: not implemented"; return false }

func nextByte(s string, offset int) (byte, int) { _ = "STUB: not implemented"; return 0, 0 }

type SPF struct {
	Hdr RR_Header
	Txt []string `dns:"txt"`
}

func (rr *SPF) String() string { _ = "STUB: not implemented"; return "" }

type AVC struct {
	Hdr RR_Header
	Txt []string `dns:"txt"`
}

func (rr *AVC) String() string { _ = "STUB: not implemented"; return "" }

type SRV struct {
	Hdr      RR_Header
	Priority uint16
	Weight   uint16
	Port     uint16
	Target   string `dns:"domain-name"`
}

func (rr *SRV) String() string { _ = "STUB: not implemented"; return "" }

type NAPTR struct {
	Hdr         RR_Header
	Order       uint16
	Preference  uint16
	Flags       string
	Service     string
	Regexp      string
	Replacement string `dns:"domain-name"`
}

func (rr *NAPTR) String() string { _ = "STUB: not implemented"; return "" }

type CERT struct {
	Hdr         RR_Header
	Type        uint16
	KeyTag      uint16
	Algorithm   uint8
	Certificate string `dns:"base64"`
}

func (rr *CERT) String() string { _ = "STUB: not implemented"; return "" }

type DNAME struct {
	Hdr    RR_Header
	Target string `dns:"domain-name"`
}

func (rr *DNAME) String() string { _ = "STUB: not implemented"; return "" }

type A struct {
	Hdr RR_Header
	A   net.IP `dns:"a"`
}

func (rr *A) String() string { _ = "STUB: not implemented"; return "" }

type AAAA struct {
	Hdr  RR_Header
	AAAA net.IP `dns:"aaaa"`
}

func (rr *AAAA) String() string { _ = "STUB: not implemented"; return "" }

type PX struct {
	Hdr        RR_Header
	Preference uint16
	Map822     string `dns:"domain-name"`
	Mapx400    string `dns:"domain-name"`
}

func (rr *PX) String() string { _ = "STUB: not implemented"; return "" }

type GPOS struct {
	Hdr       RR_Header
	Longitude string
	Latitude  string
	Altitude  string
}

func (rr *GPOS) String() string { _ = "STUB: not implemented"; return "" }

type LOC struct {
	Hdr       RR_Header
	Version   uint8
	Size      uint8
	HorizPre  uint8
	VertPre   uint8
	Latitude  uint32
	Longitude uint32
	Altitude  uint32
}

func cmToM(x uint8) string { _ = "STUB: not implemented"; return "" }

func (rr *LOC) String() string { _ = "STUB: not implemented"; return "" }

type SIG struct {
	RRSIG
}

type RRSIG struct {
	Hdr         RR_Header
	TypeCovered uint16
	Algorithm   uint8
	Labels      uint8
	OrigTtl     uint32
	Expiration  uint32
	Inception   uint32
	KeyTag      uint16
	SignerName  string `dns:"domain-name"`
	Signature   string `dns:"base64"`
}

func (rr *RRSIG) String() string { _ = "STUB: not implemented"; return "" }

type NXT struct {
	NSEC
}

type NSEC struct {
	Hdr        RR_Header
	NextDomain string   `dns:"domain-name"`
	TypeBitMap []uint16 `dns:"nsec"`
}

func (rr *NSEC) String() string { _ = "STUB: not implemented"; return "" }

func (rr *NSEC) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

type DLV struct{ DS }

type CDS struct{ DS }

type DS struct {
	Hdr        RR_Header
	KeyTag     uint16
	Algorithm  uint8
	DigestType uint8
	Digest     string `dns:"hex"`
}

func (rr *DS) String() string { _ = "STUB: not implemented"; return "" }

type KX struct {
	Hdr        RR_Header
	Preference uint16
	Exchanger  string `dns:"domain-name"`
}

func (rr *KX) String() string { _ = "STUB: not implemented"; return "" }

type TA struct {
	Hdr        RR_Header
	KeyTag     uint16
	Algorithm  uint8
	DigestType uint8
	Digest     string `dns:"hex"`
}

func (rr *TA) String() string { _ = "STUB: not implemented"; return "" }

type TALINK struct {
	Hdr          RR_Header
	PreviousName string `dns:"domain-name"`
	NextName     string `dns:"domain-name"`
}

func (rr *TALINK) String() string { _ = "STUB: not implemented"; return "" }

type SSHFP struct {
	Hdr         RR_Header
	Algorithm   uint8
	Type        uint8
	FingerPrint string `dns:"hex"`
}

func (rr *SSHFP) String() string { _ = "STUB: not implemented"; return "" }

type KEY struct {
	DNSKEY
}

type CDNSKEY struct {
	DNSKEY
}

type DNSKEY struct {
	Hdr       RR_Header
	Flags     uint16
	Protocol  uint8
	Algorithm uint8
	PublicKey string `dns:"base64"`
}

func (rr *DNSKEY) String() string { _ = "STUB: not implemented"; return "" }

type IPSECKEY struct {
	Hdr         RR_Header
	Precedence  uint8
	GatewayType uint8
	Algorithm   uint8
	GatewayAddr net.IP `dns:"-"`
	GatewayHost string `dns:"ipsechost"`
	PublicKey   string `dns:"base64"`
}

func (rr *IPSECKEY) String() string { _ = "STUB: not implemented"; return "" }

type AMTRELAY struct {
	Hdr         RR_Header
	Precedence  uint8
	GatewayType uint8
	GatewayAddr net.IP `dns:"-"`
	GatewayHost string `dns:"amtrelayhost"`
}

func (rr *AMTRELAY) String() string { _ = "STUB: not implemented"; return "" }

type RKEY struct {
	Hdr       RR_Header
	Flags     uint16
	Protocol  uint8
	Algorithm uint8
	PublicKey string `dns:"base64"`
}

func (rr *RKEY) String() string { _ = "STUB: not implemented"; return "" }

type NSAPPTR struct {
	Hdr RR_Header
	Ptr string `dns:"domain-name"`
}

func (rr *NSAPPTR) String() string { _ = "STUB: not implemented"; return "" }

type NSEC3 struct {
	Hdr        RR_Header
	Hash       uint8
	Flags      uint8
	Iterations uint16
	SaltLength uint8
	Salt       string `dns:"size-hex:SaltLength"`
	HashLength uint8
	NextDomain string   `dns:"size-base32:HashLength"`
	TypeBitMap []uint16 `dns:"nsec"`
}

func (rr *NSEC3) String() string { _ = "STUB: not implemented"; return "" }

func (rr *NSEC3) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

type NSEC3PARAM struct {
	Hdr        RR_Header
	Hash       uint8
	Flags      uint8
	Iterations uint16
	SaltLength uint8
	Salt       string `dns:"size-hex:SaltLength"`
}

func (rr *NSEC3PARAM) String() string { _ = "STUB: not implemented"; return "" }

type TKEY struct {
	Hdr        RR_Header
	Algorithm  string `dns:"domain-name"`
	Inception  uint32
	Expiration uint32
	Mode       uint16
	Error      uint16
	KeySize    uint16
	Key        string `dns:"size-hex:KeySize"`
	OtherLen   uint16
	OtherData  string `dns:"size-hex:OtherLen"`
}

func (rr *TKEY) String() string { _ = "STUB: not implemented"; return "" }

type RFC3597 struct {
	Hdr   RR_Header
	Rdata string `dns:"hex"`
}

func (rr *RFC3597) String() string { _ = "STUB: not implemented"; return "" }

func rfc3597Header(h RR_Header) string { _ = "STUB: not implemented"; return "" }

type URI struct {
	Hdr      RR_Header
	Priority uint16
	Weight   uint16
	Target   string `dns:"octet"`
}

func (rr *URI) String() string { _ = "STUB: not implemented"; return "" }

type DHCID struct {
	Hdr    RR_Header
	Digest string `dns:"base64"`
}

func (rr *DHCID) String() string { _ = "STUB: not implemented"; return "" }

type TLSA struct {
	Hdr          RR_Header
	Usage        uint8
	Selector     uint8
	MatchingType uint8
	Certificate  string `dns:"hex"`
}

func (rr *TLSA) String() string { _ = "STUB: not implemented"; return "" }

type SMIMEA struct {
	Hdr          RR_Header
	Usage        uint8
	Selector     uint8
	MatchingType uint8
	Certificate  string `dns:"hex"`
}

func (rr *SMIMEA) String() string { _ = "STUB: not implemented"; return "" }

type HIP struct {
	Hdr                RR_Header
	HitLength          uint8
	PublicKeyAlgorithm uint8
	PublicKeyLength    uint16
	Hit                string   `dns:"size-hex:HitLength"`
	PublicKey          string   `dns:"size-base64:PublicKeyLength"`
	RendezvousServers  []string `dns:"domain-name"`
}

func (rr *HIP) String() string { _ = "STUB: not implemented"; return "" }

type NINFO struct {
	Hdr    RR_Header
	ZSData []string `dns:"txt"`
}

func (rr *NINFO) String() string { _ = "STUB: not implemented"; return "" }

type NID struct {
	Hdr        RR_Header
	Preference uint16
	NodeID     uint64
}

func (rr *NID) String() string { _ = "STUB: not implemented"; return "" }

type L32 struct {
	Hdr        RR_Header
	Preference uint16
	Locator32  net.IP `dns:"a"`
}

func (rr *L32) String() string { _ = "STUB: not implemented"; return "" }

type L64 struct {
	Hdr        RR_Header
	Preference uint16
	Locator64  uint64
}

func (rr *L64) String() string { _ = "STUB: not implemented"; return "" }

type LP struct {
	Hdr        RR_Header
	Preference uint16
	Fqdn       string `dns:"domain-name"`
}

func (rr *LP) String() string { _ = "STUB: not implemented"; return "" }

type EUI48 struct {
	Hdr     RR_Header
	Address uint64 `dns:"uint48"`
}

func (rr *EUI48) String() string { _ = "STUB: not implemented"; return "" }

type EUI64 struct {
	Hdr     RR_Header
	Address uint64
}

func (rr *EUI64) String() string { _ = "STUB: not implemented"; return "" }

type CAA struct {
	Hdr   RR_Header
	Flag  uint8
	Tag   string
	Value string `dns:"octet"`
}

func (rr *CAA) String() string { _ = "STUB: not implemented"; return "" }

type UID struct {
	Hdr RR_Header
	Uid uint32
}

func (rr *UID) String() string { _ = "STUB: not implemented"; return "" }

type GID struct {
	Hdr RR_Header
	Gid uint32
}

func (rr *GID) String() string { _ = "STUB: not implemented"; return "" }

type UINFO struct {
	Hdr   RR_Header
	Uinfo string
}

func (rr *UINFO) String() string { _ = "STUB: not implemented"; return "" }

type EID struct {
	Hdr      RR_Header
	Endpoint string `dns:"hex"`
}

func (rr *EID) String() string { _ = "STUB: not implemented"; return "" }

type NIMLOC struct {
	Hdr     RR_Header
	Locator string `dns:"hex"`
}

func (rr *NIMLOC) String() string { _ = "STUB: not implemented"; return "" }

type OPENPGPKEY struct {
	Hdr       RR_Header
	PublicKey string `dns:"base64"`
}

func (rr *OPENPGPKEY) String() string { _ = "STUB: not implemented"; return "" }

type CSYNC struct {
	Hdr        RR_Header
	Serial     uint32
	Flags      uint16
	TypeBitMap []uint16 `dns:"nsec"`
}

func (rr *CSYNC) String() string { _ = "STUB: not implemented"; return "" }

func (rr *CSYNC) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

type ZONEMD struct {
	Hdr    RR_Header
	Serial uint32
	Scheme uint8
	Hash   uint8
	Digest string `dns:"hex"`
}

func (rr *ZONEMD) String() string { _ = "STUB: not implemented"; return "" }

type RESINFO struct {
	Hdr RR_Header
	Txt []string `dns:"txt"`
}

func (rr *RESINFO) String() string { _ = "STUB: not implemented"; return "" }

type APL struct {
	Hdr      RR_Header
	Prefixes []APLPrefix `dns:"apl"`
}

type APLPrefix struct {
	Negation bool
	Network  net.IPNet
}

func (rr *APL) String() string { _ = "STUB: not implemented"; return "" }

func (a *APLPrefix) str() string { _ = "STUB: not implemented"; return "" }

func (a *APLPrefix) equals(b *APLPrefix) bool { _ = "STUB: not implemented"; return false }

func (a *APLPrefix) copy() APLPrefix { _ = "STUB: not implemented"; return *new(APLPrefix) }

func (a *APLPrefix) len() int { _ = "STUB: not implemented"; return 0 }

func TimeToString(t uint32) string { _ = "STUB: not implemented"; return "" }

func StringToTime(s string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func saltToString(s string) string { _ = "STUB: not implemented"; return "" }

func euiToString(eui uint64, bits int) (hex string) { _ = "STUB: not implemented"; return "" }

func cloneSlice[E any, S ~[]E](s S) S { _ = "STUB: not implemented"; return *new(S) }

func copyNet(n net.IPNet) net.IPNet { _ = "STUB: not implemented"; return *new(net.IPNet) }

func splitN(s string, n int) []string { _ = "STUB: not implemented"; return nil }
