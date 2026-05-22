package message

import (
	"pandora/internal/pkg"
)

type AccountBannedMessage struct {
	pkg.OutboundMessage
}

func NewAccountBannedMessage() AccountBannedMessage {
	return AccountBannedMessage{}
}

func (m AccountBannedMessage) GetHeader() string {
	return "AlEb"
}

func (m AccountBannedMessage) Serialize() (string, error) {
	return "", nil
}
