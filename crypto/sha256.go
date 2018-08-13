package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 .
func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
