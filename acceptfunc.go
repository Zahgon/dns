package dns

type MsgAcceptFunc func(dh Header) MsgAcceptAction

var DefaultMsgAcceptFunc MsgAcceptFunc = defaultMsgAcceptFunc

type MsgAcceptAction int

const (
	MsgAccept MsgAcceptAction = iota
	MsgReject
	MsgIgnore
	MsgRejectNotImplemented
)

func defaultMsgAcceptFunc(dh Header) MsgAcceptAction {
	_ = "STUB: not implemented"
	return *new(MsgAcceptAction)
}
