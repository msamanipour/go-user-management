package crypto_utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func Base64Encode(val string) string {
	return base64.StdEncoding.EncodeToString([]byte(val))
}

func Base64Decode(val string) string {
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return ""
	}
	return string(data)
}
