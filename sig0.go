package dns

import (
	"crypto"
)

func (rr *SIG) Sign(k crypto.Signer, m *Msg) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rr *SIG) Verify(k *KEY, buf []byte) error { _ = "STUB: not implemented"; return nil }
