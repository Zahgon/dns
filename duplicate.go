package dns

//go:generate go run duplicate_generate.go

func IsDuplicate(r1, r2 RR) bool { _ = "STUB: not implemented"; return false }

func (r1 *RR_Header) isDuplicate(_r2 RR) bool { _ = "STUB: not implemented"; return false }

func isDuplicateName(s1, s2 string) bool { _ = "STUB: not implemented"; return false }
