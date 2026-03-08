package pkg

import (
	"crypto/rand"
	"encoding/base64"
)

type SecureTokenGenerator struct {
	Size int
}

func (g SecureTokenGenerator) New() (string, error) {
	size := g.Size
	if size <= 0 {
		size = 32
	}

	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
