package common

import (
	"encoding/base32"
)

func ConvertUrlToBase48(url string) string {
	encoded := base32.StdEncoding.EncodeToString([]byte(url))
	return encoded
}

func VerifyUrlHash(url, hash string) bool {
	h := base32.StdEncoding.EncodeToString([]byte(url))
	return h == hash
}
