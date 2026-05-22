package message

import (
	"fmt"
	"pandora/internal/pkg"
	"strings"
)

type ClientVersionMessage struct {
	pkg.InboundMessage
	Version string
	Lang    string
}

func NewClientVersionMessage(message string) (ClientVersionMessage, error) {
	parts := strings.SplitN(message, "|", 2)
	if len(parts) != 2 {
		return ClientVersionMessage{}, fmt.Errorf("malformed message %s", message)
	}

	version := strings.TrimSuffix(strings.TrimSpace(parts[0]), "e")
	if len(version) == 0 {
		return ClientVersionMessage{}, fmt.Errorf("malformed message %s", message)
	}

	lang := strings.TrimSpace(parts[1])

	return ClientVersionMessage{
		Version: version,
		Lang:    lang,
	}, nil
}
