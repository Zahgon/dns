package dns

import (
	"io"
)

type ClientConfig struct {
	Servers  []string
	Search   []string
	Port     string
	Ndots    int
	Timeout  int
	Attempts int
}

func ClientConfigFromFile(resolvconf string) (*ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ClientConfigFromReader(resolvconf io.Reader) (*ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConfig) NameList(name string) []string { _ = "STUB: not implemented"; return nil }
