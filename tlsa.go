package dns

import (
	"crypto/x509"
)

func (r *TLSA) Sign(usage, selector, matchingType int, cert *x509.Certificate) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *TLSA) Verify(cert *x509.Certificate) error { _ = "STUB: not implemented"; return nil }

func TLSAName(name, service, network string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
