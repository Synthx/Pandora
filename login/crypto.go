package login

import (
	"crypto/rand"
	"fmt"
)

const (
	Lower    = "abcdefghijklmnopqrstuvwxyz"
	LowerLen = len(Lower)
)

func RandomSalt(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("invalid salt length: %d", n)
	}

	p, err := randomBytes(n)
	if err != nil {
		return "", err
	}

	for i, v := range p {
		p[i] = Lower[v%byte(LowerLen)]
	}

	return string(p), nil
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
