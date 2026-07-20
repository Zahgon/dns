package dns

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"time"
)

const (
	_ uint8 = iota
	RSAMD5
	DH
	DSA
	_
	RSASHA1
	DSANSEC3SHA1
	RSASHA1NSEC3SHA1
	RSASHA256
	_
	RSASHA512
	_
	ECCGOST
	ECDSAP256SHA256
	ECDSAP384SHA384
	ED25519
	ED448
	INDIRECT   uint8 = 252
	PRIVATEDNS uint8 = 253
	PRIVATEOID uint8 = 254
)

var AlgorithmToString = map[uint8]string{
	RSAMD5:           "RSAMD5",
	DH:               "DH",
	DSA:              "DSA",
	RSASHA1:          "RSASHA1",
	DSANSEC3SHA1:     "DSA-NSEC3-SHA1",
	RSASHA1NSEC3SHA1: "RSASHA1-NSEC3-SHA1",
	RSASHA256:        "RSASHA256",
	RSASHA512:        "RSASHA512",
	ECCGOST:          "ECC-GOST",
	ECDSAP256SHA256:  "ECDSAP256SHA256",
	ECDSAP384SHA384:  "ECDSAP384SHA384",
	ED25519:          "ED25519",
	ED448:            "ED448",
	INDIRECT:         "INDIRECT",
	PRIVATEDNS:       "PRIVATEDNS",
	PRIVATEOID:       "PRIVATEOID",
}

var AlgorithmToHash = map[uint8]crypto.Hash{
	RSAMD5:           crypto.MD5,
	DSA:              crypto.SHA1,
	RSASHA1:          crypto.SHA1,
	RSASHA1NSEC3SHA1: crypto.SHA1,
	RSASHA256:        crypto.SHA256,
	ECDSAP256SHA256:  crypto.SHA256,
	ECDSAP384SHA384:  crypto.SHA384,
	RSASHA512:        crypto.SHA512,
	ED25519:          0,
}

const (
	_ uint8 = iota
	SHA1
	SHA256
	GOST94
	SHA384
	SHA512
)

var HashToString = map[uint8]string{
	SHA1:   "SHA1",
	SHA256: "SHA256",
	GOST94: "GOST94",
	SHA384: "SHA384",
	SHA512: "SHA512",
}

const (
	SEP    = 1
	REVOKE = 1 << 7
	ZONE   = 1 << 8
)

type rrsigWireFmt struct {
	TypeCovered uint16
	Algorithm   uint8
	Labels      uint8
	OrigTtl     uint32
	Expiration  uint32
	Inception   uint32
	KeyTag      uint16
	SignerName  string `dns:"domain-name"`
}

type dnskeyWireFmt struct {
	Flags     uint16
	Protocol  uint8
	Algorithm uint8
	PublicKey string `dns:"base64"`
}

func (k *DNSKEY) KeyTag() uint16 { _ = "STUB: not implemented"; return 0 }

func (k *DNSKEY) ToDS(h uint8) *DS { _ = "STUB: not implemented"; return nil }

func (k *DNSKEY) ToCDNSKEY() *CDNSKEY { _ = "STUB: not implemented"; return nil }

func (d *DS) ToCDS() *CDS { _ = "STUB: not implemented"; return nil }

func (rr *RRSIG) Sign(k crypto.Signer, rrset []RR) error { _ = "STUB: not implemented"; return nil }

func (rr *RRSIG) signAsIs(k crypto.Signer, rrset []RR) error { _ = "STUB: not implemented"; return nil }

func sign(k crypto.Signer, hashed []byte, hash crypto.Hash, alg uint8) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rr *RRSIG) Verify(k *DNSKEY, rrset []RR) error { _ = "STUB: not implemented"; return nil }

func (rr *RRSIG) ValidityPeriod(t time.Time) bool { _ = "STUB: not implemented"; return false }

func (rr *RRSIG) sigBuf() []byte { _ = "STUB: not implemented"; return nil }

func (k *DNSKEY) publicKeyRSA() *rsa.PublicKey { _ = "STUB: not implemented"; return nil }

func (k *DNSKEY) publicKeyECDSA() *ecdsa.PublicKey { _ = "STUB: not implemented"; return nil }

func (k *DNSKEY) publicKeyED25519() ed25519.PublicKey {
	_ = "STUB: not implemented"
	return *new(ed25519.PublicKey)
}

type wireSlice [][]byte

func (p wireSlice) Len() int           { _ = "STUB: not implemented"; return 0 }
func (p wireSlice) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (p wireSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func rawSignatureData(rrset []RR, s *RRSIG) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func packSigWire(sw *rrsigWireFmt, msg []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packKeyWire(dw *dnskeyWireFmt, msg []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
