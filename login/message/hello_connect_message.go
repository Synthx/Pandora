package login_message

type HelloConnectMessage struct {
	OutboundMessage
	Salt string
}

func NewHelloConnectMessage(salt string) *HelloConnectMessage {
	message := &HelloConnectMessage{
		Salt: salt,
	}

	return message
}

func (m *HelloConnectMessage) GetHeader() string {
	return "HC"
}

func (m *HelloConnectMessage) Serialize() (string, error) {
	return m.Salt, nil
}
