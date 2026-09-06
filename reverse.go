package dns

var StringToType = reverseInt16(TypeToString)

var StringToClass = reverseInt16(ClassToString)

var StringToOpcode = reverseInt(OpcodeToString)

var StringToRcode = reverseInt(RcodeToString)

func init() {

	StringToRcode["NOTIMPL"] = RcodeNotImplemented
}

var StringToAlgorithm = reverseInt8(AlgorithmToString)

var StringToHash = reverseInt8(HashToString)

var StringToCertType = reverseInt16(CertTypeToString)

var StringToStatefulType = reverseInt16(StatefulTypeToString)

func reverseInt8(m map[uint8]string) map[string]uint8 { _ = "STUB: not implemented"; return nil }

func reverseInt16(m map[uint16]string) map[string]uint16 { _ = "STUB: not implemented"; return nil }

func reverseInt(m map[int]string) map[string]int { _ = "STUB: not implemented"; return nil }
