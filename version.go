package dns

var Version = v{1, 1, 72}

type v struct {
	Major, Minor, Patch int
}

func (v v) String() string { _ = "STUB: not implemented"; return "" }
