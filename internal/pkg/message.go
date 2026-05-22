package pkg

type BaseMessage interface {
	GetHeader() string
}

type OutboundMessage interface {
	BaseMessage
	Serialize() (string, error)
}

type InboundMessage interface {
	BaseMessage
}

type BidirectionalMessage interface {
	OutboundMessage
	InboundMessage
}
