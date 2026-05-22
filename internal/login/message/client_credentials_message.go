package message

import (
	"fmt"
	"pandora/internal/pkg"
	"strings"
)

type ClientCredentialsMessage struct {
	pkg.InboundMessage
	account string
	token   string
}

func NewClientCredentialsMessage(message string) (ClientCredentialsMessage, error) {
	parts := strings.Split(message, "\n#")
	if len(parts) != 2 {
		return ClientCredentialsMessage{}, fmt.Errorf("malformed message %s", message)
	}

	return ClientCredentialsMessage{
		account: parts[0],
		token:   parts[1],
	}, nil
}
