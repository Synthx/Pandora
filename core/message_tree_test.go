package core

import (
	"testing"
)

type mockInbound struct {
	header string
}

func (m *mockInbound) GetHeader() string { return m.header }

func TestMessageTree(t *testing.T) {
	tree := NewMessageTree()

	// Register multiple headers with different lengths
	tree.Register("A", func(payload string) (any, error) {
		return &mockInbound{header: "A"}, nil
	})
	tree.Register("Af", func(payload string) (any, error) {
		return &mockInbound{header: "Af"}, nil
	})
	tree.Register("AlEb", func(payload string) (any, error) {
		return &mockInbound{header: "AlEb"}, nil
	})

	tests := []struct {
		input    string
		expected string
	}{
		{"Af0|0", "Af"},
		{"AlEbPayload", "AlEb"},
		{"AX", "A"}, // Should match 'A' since 'AX' is not registered
	}

	for _, tt := range tests {
		msg, err := tree.Parse(tt.input)
		if err != nil {
			t.Errorf("Parse(%s) error: %v", tt.input, err)
			continue
		}

		actualHeader := msg.(*mockInbound).GetHeader()
		if actualHeader != tt.expected {
			t.Errorf("Parse(%s) = %s; want %s", tt.input, actualHeader, tt.expected)
		}
	}
}
