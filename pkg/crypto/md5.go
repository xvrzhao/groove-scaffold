package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 Calculate the md5 hash of str and return the hex string of the result.
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
