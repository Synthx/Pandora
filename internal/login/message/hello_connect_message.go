package message

import (
	"pandora/internal/pkg"
)

type HelloConnectMessage struct {
	pkg.OutboundMessage
	Salt string
}

func NewHelloConnectMessage(salt string) HelloConnectMessage {
	return HelloConnectMessage{
		Salt: salt,
	}
}

func (m HelloConnectMessage) GetHeader() string {
	return "HC"
}

func (m HelloConnectMessage) Serialize() (string, error) {
	return m.Salt, nil
}
