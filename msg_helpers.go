package dns

import (
	"encoding/base32"
	"net"
)

func unpackDataA(msg []byte, off int) (net.IP, int, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), 0, nil
}

func packDataA(a net.IP, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackDataAAAA(msg []byte, off int) (net.IP, int, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), 0, nil
}

func packDataAAAA(aaaa net.IP, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackHeader(msg []byte, off int) (rr RR_Header, off1 int, truncmsg []byte, err error) {
	_ = "STUB: not implemented"
	return *new(RR_Header), 0, nil, nil
}

func (hdr RR_Header) packHeader(msg []byte, off int, compression compressionMap, compress bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func truncateMsgFromRdlength(msg []byte, off int, rdlength uint16) (truncmsg []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var base32HexNoPadEncoding = base32.HexEncoding.WithPadding(base32.NoPadding)

func fromBase32(s []byte) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func toBase32(b []byte) string { _ = "STUB: not implemented"; return "" }

func fromBase64(s []byte) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func toBase64(b []byte) string { _ = "STUB: not implemented"; return "" }

func noRdata(h RR_Header) bool { _ = "STUB: not implemented"; return false }

func unpackUint8(msg []byte, off int) (i uint8, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func packUint8(i uint8, msg []byte, off int) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackUint16(msg []byte, off int) (i uint16, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func packUint16(i uint16, msg []byte, off int) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackUint32(msg []byte, off int) (i uint32, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func packUint32(i uint32, msg []byte, off int) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackUint48(msg []byte, off int) (i uint64, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func packUint48(i uint64, msg []byte, off int) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackUint64(msg []byte, off int) (i uint64, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func packUint64(i uint64, msg []byte, off int) (off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackString(msg []byte, off int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packString(s string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackStringBase32(msg []byte, off, end int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packStringBase32(s string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackStringBase64(msg []byte, off, end int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packStringBase64(s string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackStringHex(msg []byte, off, end int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packStringHex(s string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackStringAny(msg []byte, off, end int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packStringAny(s string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackStringTxt(msg []byte, off int) ([]string, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func packStringTxt(s []string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackDataOpt(msg []byte, off int) ([]EDNS0, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func packDataOpt(options []EDNS0, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackStringOctet(msg []byte, off int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func packStringOctet(s string, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackDataNsec(msg []byte, off int) ([]uint16, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func typeBitMapLen(bitmap []uint16) int { _ = "STUB: not implemented"; return 0 }

func packDataNsec(bitmap []uint16, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackDataSVCB(msg []byte, off int) ([]SVCBKeyValue, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func packDataSVCB(pairs []SVCBKeyValue, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackDataDomainNames(msg []byte, off, end int) ([]string, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func packDataDomainNames(names []string, msg []byte, off int, compression compressionMap, compress bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packDataApl(data []APLPrefix, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func packDataAplPrefix(p *APLPrefix, msg []byte, off int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func unpackDataApl(msg []byte, off int) ([]APLPrefix, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func unpackDataAplPrefix(msg []byte, off int) (APLPrefix, int, error) {
	_ = "STUB: not implemented"
	return *new(APLPrefix), 0, nil
}

func unpackIPSECGateway(msg []byte, off int, gatewayType uint8) (net.IP, string, int, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), "", 0, nil
}

func packIPSECGateway(gatewayAddr net.IP, gatewayString string, msg []byte, off int, gatewayType uint8, compression compressionMap, compress bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
