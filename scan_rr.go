package dns

import (
	"net"
)

func endingToString(c *zlexer, errstr string) (string, *ParseError) {
	_ = "STUB: not implemented"
	return "", nil
}

func endingToTxtSlice(c *zlexer, errstr string) ([]string, *ParseError) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rr *A) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *AAAA) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NS) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *PTR) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NSAPPTR) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *RP) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MR) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MB) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MG) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *HINFO) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *ISDN) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MINFO) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MF) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MD) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *MX) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *RT) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *AFSDB) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *X25) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *KX) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *CNAME) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *DNAME) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *SOA) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *SRV) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NAPTR) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *TALINK) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *LOC) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *HIP) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *CERT) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *OPENPGPKEY) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *CSYNC) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *ZONEMD) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *SIG) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *RRSIG) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NXT) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NSEC) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NSEC3) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NSEC3PARAM) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *EUI48) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *EUI64) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *SSHFP) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *DNSKEY) parseDNSKEY(c *zlexer, o, typ string) *ParseError {
	_ = "STUB: not implemented"
	return nil
}

func (rr *DNSKEY) parse(c *zlexer, o string) *ParseError  { _ = "STUB: not implemented"; return nil }
func (rr *KEY) parse(c *zlexer, o string) *ParseError     { _ = "STUB: not implemented"; return nil }
func (rr *CDNSKEY) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }
func (rr *DS) parse(c *zlexer, o string) *ParseError      { _ = "STUB: not implemented"; return nil }
func (rr *DLV) parse(c *zlexer, o string) *ParseError     { _ = "STUB: not implemented"; return nil }
func (rr *CDS) parse(c *zlexer, o string) *ParseError     { _ = "STUB: not implemented"; return nil }

func (rr *IPSECKEY) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *AMTRELAY) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func parseAddrHostUnion(token, o string, gatewayType uint8) (addr net.IP, host string, err error) {
	_ = "STUB: not implemented"
	return *new(net.IP), "", nil
}

func (rr *RKEY) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *EID) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NIMLOC) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *GPOS) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *DS) parseDS(c *zlexer, o, typ string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *TA) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *TLSA) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *SMIMEA) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *RFC3597) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *SPF) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *AVC) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *TXT) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NINFO) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *RESINFO) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *URI) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *DHCID) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *NID) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *L32) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *LP) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *L64) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *UID) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *GID) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *UINFO) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *PX) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *CAA) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *TKEY) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func (rr *APL) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func escapedStringOffset(s string, desiredByteOffset int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}
