package dns

import (
	"io"
	"io/fs"
)

const maxTok = 512

const maxIncludeDepth = 7

const (
	zEOF = iota
	zString
	zBlank
	zQuote
	zNewline
	zRrtpe
	zOwner
	zClass
	zDirOrigin
	zDirTTL
	zDirInclude
	zDirGenerate

	zValue
	zKey

	zExpectOwnerDir
	zExpectOwnerBl
	zExpectAny
	zExpectAnyNoClass
	zExpectAnyNoClassBl
	zExpectAnyNoTTL
	zExpectAnyNoTTLBl
	zExpectRrtype
	zExpectRrtypeBl
	zExpectRdata
	zExpectDirTTLBl
	zExpectDirTTL
	zExpectDirOriginBl
	zExpectDirOrigin
	zExpectDirIncludeBl
	zExpectDirInclude
	zExpectDirGenerate
	zExpectDirGenerateBl
)

type ParseError struct {
	file       string
	err        string
	wrappedErr error
	lex        lex
}

func (e *ParseError) Error() (s string) { _ = "STUB: not implemented"; return "" }

func (e *ParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type lex struct {
	token  string
	err    bool
	value  uint8
	torc   uint16
	line   int
	column int
}

type ttlState struct {
	ttl           uint32
	isByDirective bool
}

func NewRR(s string) (RR, error) { _ = "STUB: not implemented"; return *new(RR), nil }

func ReadRR(r io.Reader, file string) (RR, error) { _ = "STUB: not implemented"; return *new(RR), nil }

type ZoneParser struct {
	c *zlexer

	parseErr *ParseError

	origin string
	file   string

	defttl *ttlState

	h RR_Header

	sub  *ZoneParser
	r    io.Reader
	fsys fs.FS

	includeDepth uint8

	includeAllowed     bool
	generateDisallowed bool
}

func NewZoneParser(r io.Reader, origin, file string) *ZoneParser {
	_ = "STUB: not implemented"
	return nil
}

func (zp *ZoneParser) SetDefaultTTL(ttl uint32) { _ = "STUB: not implemented"; return }

func (zp *ZoneParser) SetIncludeAllowed(v bool) { _ = "STUB: not implemented"; return }

func (zp *ZoneParser) SetIncludeFS(fsys fs.FS) { _ = "STUB: not implemented"; return }

func (zp *ZoneParser) Err() error { _ = "STUB: not implemented"; return nil }

func (zp *ZoneParser) setParseError(err string, l lex) (RR, bool) {
	_ = "STUB: not implemented"
	return *new(RR), false
}

func (zp *ZoneParser) Comment() string { _ = "STUB: not implemented"; return "" }

func (zp *ZoneParser) subNext() (RR, bool) { _ = "STUB: not implemented"; return *new(RR), false }

func (zp *ZoneParser) Next() (RR, bool) { _ = "STUB: not implemented"; return *new(RR), false }

type zlexer struct {
	br io.ByteReader

	readErr error

	line   int
	column int

	comBuf  string
	comment string

	l       lex
	cachedL *lex

	brace  int
	quote  bool
	space  bool
	commt  bool
	rrtype bool
	owner  bool

	nextL bool

	eol bool
}

func newZLexer(r io.Reader) *zlexer { _ = "STUB: not implemented"; return nil }

func (zl *zlexer) Err() error { _ = "STUB: not implemented"; return nil }

func (zl *zlexer) readByte() (byte, bool) { _ = "STUB: not implemented"; return 0, false }

func (zl *zlexer) Peek() lex { _ = "STUB: not implemented"; return *new(lex) }

func (zl *zlexer) Next() (lex, bool) { _ = "STUB: not implemented"; return *new(lex), false }

func (zl *zlexer) Comment() string { _ = "STUB: not implemented"; return "" }

func classToInt(token string) (uint16, bool) { _ = "STUB: not implemented"; return 0, false }

func typeToInt(token string) (uint16, bool) { _ = "STUB: not implemented"; return 0, false }

func stringToTTL(token string) (uint32, bool) { _ = "STUB: not implemented"; return 0, false }

func stringToCm(token string) (e, m uint8, ok bool) { _ = "STUB: not implemented"; return 0, 0, false }

func toAbsoluteName(name, origin string) (absolute string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func appendOrigin(name, origin string) string { _ = "STUB: not implemented"; return "" }

func locCheckNorth(token string, latitude uint32) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func locCheckEast(token string, longitude uint32) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func slurpRemainder(c *zlexer) *ParseError { _ = "STUB: not implemented"; return nil }

func stringToNodeID(l lex) (uint64, *ParseError) { _ = "STUB: not implemented"; return 0, nil }
