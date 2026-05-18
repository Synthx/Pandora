package login_message

type BaseMessage interface {
	GetHeader() string
}

type OutboundMessage interface {
	BaseMessage
	Serialize() (string, error)
}
