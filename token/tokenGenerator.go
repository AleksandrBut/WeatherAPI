package token

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() string {
	bytes := make([]byte, 32)
	_, _ = rand.Read(bytes)

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
}
