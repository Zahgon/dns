package dns

import (
	"bytes"
)

func (zp *ZoneParser) generate(l lex) (RR, bool) { _ = "STUB: not implemented"; return *new(RR), false }

type generateReader struct {
	s  string
	si int

	cur   int64
	start int64
	end   int64
	step  int64

	mod bytes.Buffer

	escape bool

	eof bool

	file string
	lex  *lex
}

func (r *generateReader) parseError(msg string, end int) *ParseError {
	_ = "STUB: not implemented"
	return nil
}

func (r *generateReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *generateReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func modToPrintf(s string) (string, int64, string) { _ = "STUB: not implemented"; return "", 0, "" }
