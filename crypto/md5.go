package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Sum(bs []byte, b ...byte) string {
	h := md5.New()
	h.Write(bs)
	return hex.EncodeToString(h.Sum(b))
}
