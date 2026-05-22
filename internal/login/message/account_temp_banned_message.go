package message

import (
	"fmt"
	"pandora/internal/pkg"
	"time"
)

type AccountTempBannedMessage struct {
	pkg.OutboundMessage
	date time.Time
}

func NewAccountTempBannedMessage(date time.Time) AccountTempBannedMessage {
	return AccountTempBannedMessage{
		date: date,
	}
}

func (m AccountTempBannedMessage) GetHeader() string {
	return "AlEk"
}

func (m AccountTempBannedMessage) Serialize() (string, error) {
	if m.date.IsZero() {
		return "0|0|0", nil
	}

	duration := time.Until(m.date)
	if duration <= 0 {
		return "0|0|0", nil
	}

	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%d|%d|%d", days, hours, minutes), nil
}
