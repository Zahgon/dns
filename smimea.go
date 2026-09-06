package dns

import (
	"crypto/x509"
)

func (r *SMIMEA) Sign(usage, selector, matchingType int, cert *x509.Certificate) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *SMIMEA) Verify(cert *x509.Certificate) error { _ = "STUB: not implemented"; return nil }

func SMIMEAName(email, domain string) (string, error) { _ = "STUB: not implemented"; return "", nil }
