package util

import (
	"crypto/rand"
	"encoding/hex"
)

func NewHex(length int64) string {
	bytes := make([]byte, length/2)

	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
