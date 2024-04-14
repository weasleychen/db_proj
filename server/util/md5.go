package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Encrypt(message string) string {
	hash := md5.New()
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}

func EncryptInt64(message int64) string {
	return Encrypt(fmt.Sprint(message))
}
