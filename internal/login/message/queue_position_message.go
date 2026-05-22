package message

import (
	"pandora/internal/pkg"
)

type QueuePositionMessage struct {
	pkg.BidirectionalMessage
}

func NewQueuePositionMessage() QueuePositionMessage {
	return QueuePositionMessage{}
}

func (m QueuePositionMessage) GetHeader() string {
	return "Af"
}

func (m QueuePositionMessage) Serialize() (string, error) {
	return "1|0|1|0|0", nil
}
