package crypto

import (
	"crypto/sha512"
	"encoding/hex"
)

// Sha512 .
func Sha512(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
