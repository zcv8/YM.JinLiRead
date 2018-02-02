package common

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5加密
func EncryptionMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
