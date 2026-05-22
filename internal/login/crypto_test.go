package login

import (
	"strings"
	"testing"
)

func TestRandomSalt(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"length 8", 8, false},
		{"length 16", 16, false},
		{"length 32", 32, false},
		{"zero length", 0, true},
		{"negative length", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salt, err := RandomSalt(tt.length)

			if (err != nil) != tt.wantErr {
				t.Fatalf("RandomSalt(%d) error = %v, wantErr %v", tt.length, err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if len(salt) != tt.length {
				t.Errorf("RandomSalt(%d) returned salt of length %d, want %d", tt.length, len(salt), tt.length)
			}

			for _, char := range salt {
				if !strings.ContainsRune(Lower, char) {
					t.Errorf("RandomSalt(%d) contains invalid character: %q", tt.length, char)
				}
			}
		})
	}
}

func TestRandomSalt_Randomness(t *testing.T) {
	length := 16
	s1, err := RandomSalt(length)
	if err != nil {
		t.Fatalf("Failed to generate salt 1: %v", err)
	}

	s2, err := RandomSalt(length)
	if err != nil {
		t.Fatalf("Failed to generate salt 2: %v", err)
	}

	if s1 == s2 {
		t.Errorf("RandomSalt produced identical results: %q == %q", s1, s2)
	}
}
