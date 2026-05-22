package message

import "pandora/internal/pkg"

type ChooseNicknameMessage struct {
	pkg.OutboundMessage
}

func NewChooseNicknameMessage() ChooseNicknameMessage {
	return ChooseNicknameMessage{}
}

func (m ChooseNicknameMessage) GetHeader() string {
	return "AlEr"
}

func (m ChooseNicknameMessage) Serialize() (string, error) {
	return "", nil
}
