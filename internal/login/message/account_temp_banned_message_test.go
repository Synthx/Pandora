package message

import (
	"testing"
	"time"
)

func TestAccountTempBannedMessage(t *testing.T) {
	// Set up a ban date 2 days, 5 hours, and 30 minutes from now
	// We add a few seconds to avoid race conditions with time.Now() in Serialize
	banDate := time.Now().Add(48*time.Hour + 5*time.Hour + 30*time.Minute + 2*time.Second)

	msg := NewAccountTempBannedMessage(banDate)

	// Test Header
	expectedHeader := "AlEk"
	if msg.GetHeader() != expectedHeader {
		t.Errorf("GetHeader() = %s; want %s", msg.GetHeader(), expectedHeader)
	}

	// Test Serialization
	serialized, err := msg.Serialize()
	if err != nil {
		t.Fatalf("Serialize() error: %v", err)
	}

	// Expected: 2 days, 5 hours, 30 minutes
	expectedSerialized := "2|5|30"

	if serialized != expectedSerialized {
		t.Errorf("Serialize() = %s; want %s", serialized, expectedSerialized)
	}
}
