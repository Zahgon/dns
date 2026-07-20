package dns

func HashName(label string, ha uint8, iter uint16, salt string) string {
	_ = "STUB: not implemented"
	return ""
}

func (rr *NSEC3) Cover(name string) bool { _ = "STUB: not implemented"; return false }

func (rr *NSEC3) Match(name string) bool { _ = "STUB: not implemented"; return false }
