package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString(length int) (string, error) {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
