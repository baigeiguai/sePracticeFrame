package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Convert(password string) string {
	m := md5.New()
	m.Write([]byte(password))
	return hex.EncodeToString(m.Sum(nil))
}
