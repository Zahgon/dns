package dns

import (
	"bytes"
	"crypto"
	"hash"
)

type identityHash struct {
	b *bytes.Buffer
}

func (i identityHash) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (i identityHash) Size() int                   { _ = "STUB: not implemented"; return 0 }
func (i identityHash) BlockSize() int              { _ = "STUB: not implemented"; return 0 }
func (i identityHash) Reset()                      { _ = "STUB: not implemented"; return }
func (i identityHash) Sum(b []byte) []byte         { _ = "STUB: not implemented"; return nil }

func hashFromAlgorithm(alg uint8) (hash.Hash, crypto.Hash, error) {
	_ = "STUB: not implemented"
	return *new(hash.Hash), *new(crypto.Hash), nil
}
