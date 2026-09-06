package dns

var TypeToRR = map[uint16]func() RR{
	TypeA:          func() RR { return new(A) },
	TypeAAAA:       func() RR { return new(AAAA) },
	TypeAFSDB:      func() RR { return new(AFSDB) },
	TypeAMTRELAY:   func() RR { return new(AMTRELAY) },
	TypeANY:        func() RR { return new(ANY) },
	TypeAPL:        func() RR { return new(APL) },
	TypeAVC:        func() RR { return new(AVC) },
	TypeCAA:        func() RR { return new(CAA) },
	TypeCDNSKEY:    func() RR { return new(CDNSKEY) },
	TypeCDS:        func() RR { return new(CDS) },
	TypeCERT:       func() RR { return new(CERT) },
	TypeCNAME:      func() RR { return new(CNAME) },
	TypeCSYNC:      func() RR { return new(CSYNC) },
	TypeDHCID:      func() RR { return new(DHCID) },
	TypeDLV:        func() RR { return new(DLV) },
	TypeDNAME:      func() RR { return new(DNAME) },
	TypeDNSKEY:     func() RR { return new(DNSKEY) },
	TypeDS:         func() RR { return new(DS) },
	TypeEID:        func() RR { return new(EID) },
	TypeEUI48:      func() RR { return new(EUI48) },
	TypeEUI64:      func() RR { return new(EUI64) },
	TypeGID:        func() RR { return new(GID) },
	TypeGPOS:       func() RR { return new(GPOS) },
	TypeHINFO:      func() RR { return new(HINFO) },
	TypeHIP:        func() RR { return new(HIP) },
	TypeHTTPS:      func() RR { return new(HTTPS) },
	TypeIPSECKEY:   func() RR { return new(IPSECKEY) },
	TypeISDN:       func() RR { return new(ISDN) },
	TypeKEY:        func() RR { return new(KEY) },
	TypeKX:         func() RR { return new(KX) },
	TypeL32:        func() RR { return new(L32) },
	TypeL64:        func() RR { return new(L64) },
	TypeLOC:        func() RR { return new(LOC) },
	TypeLP:         func() RR { return new(LP) },
	TypeMB:         func() RR { return new(MB) },
	TypeMD:         func() RR { return new(MD) },
	TypeMF:         func() RR { return new(MF) },
	TypeMG:         func() RR { return new(MG) },
	TypeMINFO:      func() RR { return new(MINFO) },
	TypeMR:         func() RR { return new(MR) },
	TypeMX:         func() RR { return new(MX) },
	TypeNAPTR:      func() RR { return new(NAPTR) },
	TypeNID:        func() RR { return new(NID) },
	TypeNIMLOC:     func() RR { return new(NIMLOC) },
	TypeNINFO:      func() RR { return new(NINFO) },
	TypeNS:         func() RR { return new(NS) },
	TypeNSAPPTR:    func() RR { return new(NSAPPTR) },
	TypeNSEC:       func() RR { return new(NSEC) },
	TypeNSEC3:      func() RR { return new(NSEC3) },
	TypeNSEC3PARAM: func() RR { return new(NSEC3PARAM) },
	TypeNULL:       func() RR { return new(NULL) },
	TypeNXNAME:     func() RR { return new(NXNAME) },
	TypeNXT:        func() RR { return new(NXT) },
	TypeOPENPGPKEY: func() RR { return new(OPENPGPKEY) },
	TypeOPT:        func() RR { return new(OPT) },
	TypePTR:        func() RR { return new(PTR) },
	TypePX:         func() RR { return new(PX) },
	TypeRESINFO:    func() RR { return new(RESINFO) },
	TypeRKEY:       func() RR { return new(RKEY) },
	TypeRP:         func() RR { return new(RP) },
	TypeRRSIG:      func() RR { return new(RRSIG) },
	TypeRT:         func() RR { return new(RT) },
	TypeSIG:        func() RR { return new(SIG) },
	TypeSMIMEA:     func() RR { return new(SMIMEA) },
	TypeSOA:        func() RR { return new(SOA) },
	TypeSPF:        func() RR { return new(SPF) },
	TypeSRV:        func() RR { return new(SRV) },
	TypeSSHFP:      func() RR { return new(SSHFP) },
	TypeSVCB:       func() RR { return new(SVCB) },
	TypeTA:         func() RR { return new(TA) },
	TypeTALINK:     func() RR { return new(TALINK) },
	TypeTKEY:       func() RR { return new(TKEY) },
	TypeTLSA:       func() RR { return new(TLSA) },
	TypeTSIG:       func() RR { return new(TSIG) },
	TypeTXT:        func() RR { return new(TXT) },
	TypeUID:        func() RR { return new(UID) },
	TypeUINFO:      func() RR { return new(UINFO) },
	TypeURI:        func() RR { return new(URI) },
	TypeX25:        func() RR { return new(X25) },
	TypeZONEMD:     func() RR { return new(ZONEMD) },
}

var TypeToString = map[uint16]string{
	TypeA:          "A",
	TypeAAAA:       "AAAA",
	TypeAFSDB:      "AFSDB",
	TypeAMTRELAY:   "AMTRELAY",
	TypeANY:        "ANY",
	TypeAPL:        "APL",
	TypeATMA:       "ATMA",
	TypeAVC:        "AVC",
	TypeAXFR:       "AXFR",
	TypeCAA:        "CAA",
	TypeCDNSKEY:    "CDNSKEY",
	TypeCDS:        "CDS",
	TypeCERT:       "CERT",
	TypeCNAME:      "CNAME",
	TypeCSYNC:      "CSYNC",
	TypeDHCID:      "DHCID",
	TypeDLV:        "DLV",
	TypeDNAME:      "DNAME",
	TypeDNSKEY:     "DNSKEY",
	TypeDS:         "DS",
	TypeEID:        "EID",
	TypeEUI48:      "EUI48",
	TypeEUI64:      "EUI64",
	TypeGID:        "GID",
	TypeGPOS:       "GPOS",
	TypeHINFO:      "HINFO",
	TypeHIP:        "HIP",
	TypeHTTPS:      "HTTPS",
	TypeIPSECKEY:   "IPSECKEY",
	TypeISDN:       "ISDN",
	TypeIXFR:       "IXFR",
	TypeKEY:        "KEY",
	TypeKX:         "KX",
	TypeL32:        "L32",
	TypeL64:        "L64",
	TypeLOC:        "LOC",
	TypeLP:         "LP",
	TypeMAILA:      "MAILA",
	TypeMAILB:      "MAILB",
	TypeMB:         "MB",
	TypeMD:         "MD",
	TypeMF:         "MF",
	TypeMG:         "MG",
	TypeMINFO:      "MINFO",
	TypeMR:         "MR",
	TypeMX:         "MX",
	TypeNAPTR:      "NAPTR",
	TypeNID:        "NID",
	TypeNIMLOC:     "NIMLOC",
	TypeNINFO:      "NINFO",
	TypeNS:         "NS",
	TypeNSEC:       "NSEC",
	TypeNSEC3:      "NSEC3",
	TypeNSEC3PARAM: "NSEC3PARAM",
	TypeNULL:       "NULL",
	TypeNXNAME:     "NXNAME",
	TypeNXT:        "NXT",
	TypeNone:       "None",
	TypeOPENPGPKEY: "OPENPGPKEY",
	TypeOPT:        "OPT",
	TypePTR:        "PTR",
	TypePX:         "PX",
	TypeRESINFO:    "RESINFO",
	TypeRKEY:       "RKEY",
	TypeRP:         "RP",
	TypeRRSIG:      "RRSIG",
	TypeRT:         "RT",
	TypeReserved:   "Reserved",
	TypeSIG:        "SIG",
	TypeSMIMEA:     "SMIMEA",
	TypeSOA:        "SOA",
	TypeSPF:        "SPF",
	TypeSRV:        "SRV",
	TypeSSHFP:      "SSHFP",
	TypeSVCB:       "SVCB",
	TypeTA:         "TA",
	TypeTALINK:     "TALINK",
	TypeTKEY:       "TKEY",
	TypeTLSA:       "TLSA",
	TypeTSIG:       "TSIG",
	TypeTXT:        "TXT",
	TypeUID:        "UID",
	TypeUINFO:      "UINFO",
	TypeUNSPEC:     "UNSPEC",
	TypeURI:        "URI",
	TypeX25:        "X25",
	TypeZONEMD:     "ZONEMD",
	TypeNSAPPTR:    "NSAP-PTR",
}

func (rr *A) Header() *RR_Header          { _ = "STUB: not implemented"; return nil }
func (rr *AAAA) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *AFSDB) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *AMTRELAY) Header() *RR_Header   { _ = "STUB: not implemented"; return nil }
func (rr *ANY) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *APL) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *AVC) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *CAA) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *CDNSKEY) Header() *RR_Header    { _ = "STUB: not implemented"; return nil }
func (rr *CDS) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *CERT) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *CNAME) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *CSYNC) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *DHCID) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *DLV) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *DNAME) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *DNSKEY) Header() *RR_Header     { _ = "STUB: not implemented"; return nil }
func (rr *DS) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *EID) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *EUI48) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *EUI64) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *GID) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *GPOS) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *HINFO) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *HIP) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *HTTPS) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *IPSECKEY) Header() *RR_Header   { _ = "STUB: not implemented"; return nil }
func (rr *ISDN) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *KEY) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *KX) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *L32) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *L64) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *LOC) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *LP) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *MB) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *MD) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *MF) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *MG) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *MINFO) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *MR) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *MX) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *NAPTR) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *NID) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *NIMLOC) Header() *RR_Header     { _ = "STUB: not implemented"; return nil }
func (rr *NINFO) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *NS) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *NSAPPTR) Header() *RR_Header    { _ = "STUB: not implemented"; return nil }
func (rr *NSEC) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *NSEC3) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *NSEC3PARAM) Header() *RR_Header { _ = "STUB: not implemented"; return nil }
func (rr *NULL) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *NXNAME) Header() *RR_Header     { _ = "STUB: not implemented"; return nil }
func (rr *NXT) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *OPENPGPKEY) Header() *RR_Header { _ = "STUB: not implemented"; return nil }
func (rr *OPT) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *PTR) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *PX) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *RESINFO) Header() *RR_Header    { _ = "STUB: not implemented"; return nil }
func (rr *RFC3597) Header() *RR_Header    { _ = "STUB: not implemented"; return nil }
func (rr *RKEY) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *RP) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *RRSIG) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *RT) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *SIG) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *SMIMEA) Header() *RR_Header     { _ = "STUB: not implemented"; return nil }
func (rr *SOA) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *SPF) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *SRV) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *SSHFP) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *SVCB) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *TA) Header() *RR_Header         { _ = "STUB: not implemented"; return nil }
func (rr *TALINK) Header() *RR_Header     { _ = "STUB: not implemented"; return nil }
func (rr *TKEY) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *TLSA) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *TSIG) Header() *RR_Header       { _ = "STUB: not implemented"; return nil }
func (rr *TXT) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *UID) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *UINFO) Header() *RR_Header      { _ = "STUB: not implemented"; return nil }
func (rr *URI) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *X25) Header() *RR_Header        { _ = "STUB: not implemented"; return nil }
func (rr *ZONEMD) Header() *RR_Header     { _ = "STUB: not implemented"; return nil }

func (rr *A) len(off int, compression map[string]struct{}) int { _ = "STUB: not implemented"; return 0 }

func (rr *AAAA) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *AFSDB) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *AMTRELAY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *ANY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *APL) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *AVC) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *CAA) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *CERT) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *CNAME) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *DHCID) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *DNAME) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *DNSKEY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *DS) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *EID) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *EUI48) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *EUI64) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *GID) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *GPOS) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *HINFO) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *HIP) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *IPSECKEY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *ISDN) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *KX) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *L32) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *L64) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *LOC) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *LP) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MB) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MD) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MF) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MG) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MINFO) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MR) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *MX) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NAPTR) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NID) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NIMLOC) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NINFO) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NS) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NSAPPTR) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NSEC3PARAM) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NULL) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *NXNAME) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *OPENPGPKEY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *PTR) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *PX) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *RESINFO) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *RFC3597) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *RKEY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *RP) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *RRSIG) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *RT) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *SMIMEA) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *SOA) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *SPF) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *SRV) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *SSHFP) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *SVCB) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *TA) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *TALINK) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *TKEY) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *TLSA) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *TSIG) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *TXT) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *UID) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *UINFO) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *URI) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *X25) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *ZONEMD) len(off int, compression map[string]struct{}) int {
	_ = "STUB: not implemented"
	return 0
}

func (rr *A) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *AAAA) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *AFSDB) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *AMTRELAY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *ANY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *APL) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *AVC) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *CAA) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *CDNSKEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *CDS) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *CERT) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *CNAME) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *CSYNC) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *DHCID) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *DLV) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *DNAME) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *DNSKEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *DS) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *EID) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *EUI48) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *EUI64) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *GID) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *GPOS) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *HINFO) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *HIP) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *HTTPS) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *IPSECKEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *ISDN) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *KEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *KX) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *L32) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *L64) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *LOC) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *LP) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MB) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MD) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MF) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MG) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MINFO) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MR) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *MX) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NAPTR) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NID) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NIMLOC) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NINFO) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NS) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NSAPPTR) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NSEC) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NSEC3) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NSEC3PARAM) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NULL) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NXNAME) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *NXT) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *OPENPGPKEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *OPT) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *PTR) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *PX) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *RESINFO) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *RFC3597) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *RKEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *RP) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *RRSIG) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *RT) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SIG) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SMIMEA) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SOA) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SPF) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SRV) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SSHFP) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *SVCB) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *TA) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *TALINK) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *TKEY) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *TLSA) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *TSIG) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *TXT) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *UID) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *UINFO) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *URI) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *X25) copy() RR { _ = "STUB: not implemented"; return *new(RR) }

func (rr *ZONEMD) copy() RR { _ = "STUB: not implemented"; return *new(RR) }
