package message

import (
	"pandora/internal/pkg"
)

type InvalidVersionMessage struct {
	pkg.OutboundMessage
	Version string
}

func NewInvalidVersionMessage(version string) InvalidVersionMessage {
	return InvalidVersionMessage{
		Version: version,
	}
}

func (m InvalidVersionMessage) GetHeader() string {
	return "AlEv"
}

func (m InvalidVersionMessage) Serialize() (string, error) {
	return m.Version, nil
}
