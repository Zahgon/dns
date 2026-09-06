package dns

import (
	"net"
)

type SVCBKey uint16

const (
	SVCB_MANDATORY SVCBKey = iota
	SVCB_ALPN
	SVCB_NO_DEFAULT_ALPN
	SVCB_PORT
	SVCB_IPV4HINT
	SVCB_ECHCONFIG
	SVCB_IPV6HINT
	SVCB_DOHPATH
	SVCB_OHTTP

	svcb_RESERVED SVCBKey = 65535
)

var svcbKeyToStringMap = map[SVCBKey]string{
	SVCB_MANDATORY:       "mandatory",
	SVCB_ALPN:            "alpn",
	SVCB_NO_DEFAULT_ALPN: "no-default-alpn",
	SVCB_PORT:            "port",
	SVCB_IPV4HINT:        "ipv4hint",
	SVCB_ECHCONFIG:       "ech",
	SVCB_IPV6HINT:        "ipv6hint",
	SVCB_DOHPATH:         "dohpath",
	SVCB_OHTTP:           "ohttp",
}

var svcbStringToKeyMap = reverseSVCBKeyMap(svcbKeyToStringMap)

func reverseSVCBKeyMap(m map[SVCBKey]string) map[string]SVCBKey {
	_ = "STUB: not implemented"
	return nil
}

func (key SVCBKey) String() string { _ = "STUB: not implemented"; return "" }

func svcbStringToKey(s string) SVCBKey { _ = "STUB: not implemented"; return *new(SVCBKey) }

func (rr *SVCB) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

func makeSVCBKeyValue(key SVCBKey) SVCBKeyValue {
	_ = "STUB: not implemented"
	return *new(SVCBKeyValue)
}

type SVCB struct {
	Hdr      RR_Header
	Priority uint16
	Target   string         `dns:"domain-name"`
	Value    []SVCBKeyValue `dns:"pairs"`
}

type HTTPS struct {
	SVCB
}

func (rr *HTTPS) String() string { _ = "STUB: not implemented"; return "" }

func (rr *HTTPS) parse(c *zlexer, o string) *ParseError { _ = "STUB: not implemented"; return nil }

type SVCBKeyValue interface {
	Key() SVCBKey
	pack() ([]byte, error)
	unpack([]byte) error
	String() string
	parse(string) error
	copy() SVCBKeyValue
	len() int
}

type SVCBMandatory struct {
	Code []SVCBKey
}

func (*SVCBMandatory) Key() SVCBKey { _ = "STUB: not implemented"; return *new(SVCBKey) }

func (s *SVCBMandatory) String() string { _ = "STUB: not implemented"; return "" }

func (s *SVCBMandatory) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBMandatory) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBMandatory) parse(b string) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBMandatory) len() int { _ = "STUB: not implemented"; return 0 }

func (s *SVCBMandatory) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

type SVCBAlpn struct {
	Alpn []string
}

func (*SVCBAlpn) Key() SVCBKey { _ = "STUB: not implemented"; return *new(SVCBKey) }

func (s *SVCBAlpn) String() string { _ = "STUB: not implemented"; return "" }

func (s *SVCBAlpn) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBAlpn) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBAlpn) parse(b string) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBAlpn) len() int { _ = "STUB: not implemented"; return 0 }

func (s *SVCBAlpn) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

type SVCBNoDefaultAlpn struct{}

func (*SVCBNoDefaultAlpn) Key() SVCBKey          { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (*SVCBNoDefaultAlpn) copy() SVCBKeyValue    { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }
func (*SVCBNoDefaultAlpn) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (*SVCBNoDefaultAlpn) String() string        { _ = "STUB: not implemented"; return "" }
func (*SVCBNoDefaultAlpn) len() int              { _ = "STUB: not implemented"; return 0 }

func (*SVCBNoDefaultAlpn) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (*SVCBNoDefaultAlpn) parse(b string) error { _ = "STUB: not implemented"; return nil }

type SVCBPort struct {
	Port uint16
}

func (*SVCBPort) Key() SVCBKey         { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (*SVCBPort) len() int             { _ = "STUB: not implemented"; return 0 }
func (s *SVCBPort) String() string     { _ = "STUB: not implemented"; return "" }
func (s *SVCBPort) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

func (s *SVCBPort) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBPort) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBPort) parse(b string) error { _ = "STUB: not implemented"; return nil }

type SVCBIPv4Hint struct {
	Hint []net.IP
}

func (*SVCBIPv4Hint) Key() SVCBKey { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (s *SVCBIPv4Hint) len() int   { _ = "STUB: not implemented"; return 0 }

func (s *SVCBIPv4Hint) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBIPv4Hint) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBIPv4Hint) String() string { _ = "STUB: not implemented"; return "" }

func (s *SVCBIPv4Hint) parse(b string) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBIPv4Hint) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

type SVCBECHConfig struct {
	ECH []byte
}

func (*SVCBECHConfig) Key() SVCBKey     { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (s *SVCBECHConfig) String() string { _ = "STUB: not implemented"; return "" }
func (s *SVCBECHConfig) len() int       { _ = "STUB: not implemented"; return 0 }

func (s *SVCBECHConfig) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBECHConfig) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

func (s *SVCBECHConfig) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBECHConfig) parse(b string) error { _ = "STUB: not implemented"; return nil }

type SVCBIPv6Hint struct {
	Hint []net.IP
}

func (*SVCBIPv6Hint) Key() SVCBKey { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (s *SVCBIPv6Hint) len() int   { _ = "STUB: not implemented"; return 0 }

func (s *SVCBIPv6Hint) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBIPv6Hint) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBIPv6Hint) String() string { _ = "STUB: not implemented"; return "" }

func (s *SVCBIPv6Hint) parse(b string) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBIPv6Hint) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

type SVCBDoHPath struct {
	Template string
}

func (*SVCBDoHPath) Key() SVCBKey            { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (s *SVCBDoHPath) String() string        { _ = "STUB: not implemented"; return "" }
func (s *SVCBDoHPath) len() int              { _ = "STUB: not implemented"; return 0 }
func (s *SVCBDoHPath) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *SVCBDoHPath) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBDoHPath) parse(b string) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBDoHPath) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

type SVCBOhttp struct{}

func (*SVCBOhttp) Key() SVCBKey          { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (*SVCBOhttp) copy() SVCBKeyValue    { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }
func (*SVCBOhttp) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (*SVCBOhttp) String() string        { _ = "STUB: not implemented"; return "" }
func (*SVCBOhttp) len() int              { _ = "STUB: not implemented"; return 0 }

func (*SVCBOhttp) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (*SVCBOhttp) parse(b string) error { _ = "STUB: not implemented"; return nil }

type SVCBLocal struct {
	KeyCode SVCBKey
	Data    []byte
}

func (s *SVCBLocal) Key() SVCBKey          { _ = "STUB: not implemented"; return *new(SVCBKey) }
func (s *SVCBLocal) String() string        { _ = "STUB: not implemented"; return "" }
func (s *SVCBLocal) pack() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (s *SVCBLocal) len() int              { _ = "STUB: not implemented"; return 0 }

func (s *SVCBLocal) unpack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBLocal) parse(b string) error { _ = "STUB: not implemented"; return nil }

func (s *SVCBLocal) copy() SVCBKeyValue { _ = "STUB: not implemented"; return *new(SVCBKeyValue) }

func (rr *SVCB) String() string { _ = "STUB: not implemented"; return "" }

func areSVCBPairArraysEqual(a []SVCBKeyValue, b []SVCBKeyValue) bool {
	_ = "STUB: not implemented"
	return false
}

func svcbParamToStr(s []byte) string { _ = "STUB: not implemented"; return "" }

func svcbParseParam(b string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
