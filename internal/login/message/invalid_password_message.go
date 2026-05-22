package message

import (
	"pandora/internal/pkg"
)

type InvalidPasswordMessage struct {
	pkg.OutboundMessage
}

func NewInvalidPasswordMessage() InvalidPasswordMessage {
	return InvalidPasswordMessage{}
}

func (m InvalidPasswordMessage) GetHeader() string {
	return "AlEf"
}

func (m InvalidPasswordMessage) Serialize() (string, error) {
	return "", nil
}
