package dns

import (
	"crypto/x509"
)

func CertificateToDANE(selector, matchingType uint8, cert *x509.Certificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
