package dns

import (
	"crypto"
	"crypto/ed25519"
	"math/big"
)

func (k *DNSKEY) Generate(bits int) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func (k *DNSKEY) setPublicKeyRSA(_E int, _N *big.Int) bool { _ = "STUB: not implemented"; return false }

func (k *DNSKEY) setPublicKeyECDSA(_X, _Y *big.Int) bool { _ = "STUB: not implemented"; return false }

func (k *DNSKEY) setPublicKeyED25519(_K ed25519.PublicKey) bool {
	_ = "STUB: not implemented"
	return false
}

func exponentToBuf(_E int) []byte { _ = "STUB: not implemented"; return nil }

func curveToBuf(_X, _Y *big.Int, intlen int) []byte { _ = "STUB: not implemented"; return nil }
