package dns

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"io"
)

func (k *DNSKEY) NewPrivateKey(s string) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func (k *DNSKEY) ReadPrivateKey(q io.Reader, file string) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func readPrivateKeyRSA(m map[string]string) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readPrivateKeyECDSA(m map[string]string) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readPrivateKeyED25519(m map[string]string) (ed25519.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(ed25519.PrivateKey), nil
}

func parseKey(r io.Reader, file string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type klexer struct {
	br io.ByteReader

	readErr error

	line   int
	column int

	key bool

	eol bool
}

func newKLexer(r io.Reader) *klexer { _ = "STUB: not implemented"; return nil }

func (kl *klexer) Err() error { _ = "STUB: not implemented"; return nil }

func (kl *klexer) readByte() (byte, bool) { _ = "STUB: not implemented"; return 0, false }

func (kl *klexer) Next() (lex, bool) { _ = "STUB: not implemented"; return *new(lex), false }
