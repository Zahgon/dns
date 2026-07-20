package dns

import (
	"crypto"
	"math/big"
)

const format = "Private-key-format: v1.3\n"

var bigIntOne = big.NewInt(1)

func (r *DNSKEY) PrivateKeyString(p crypto.PrivateKey) string { _ = "STUB: not implemented"; return "" }
