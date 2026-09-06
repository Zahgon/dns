//go:build fuzz
// +build fuzz

package dns

func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }

func FuzzNewRR(data []byte) int { _ = "STUB: not implemented"; return 0 }
