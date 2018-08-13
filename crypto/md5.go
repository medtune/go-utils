package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 .
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
