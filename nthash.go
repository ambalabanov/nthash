package nthash

import (
	"encoding/hex"
	"io"

	"golang.org/x/crypto/md4"
	"golang.org/x/text/encoding/unicode"
)

//GetHash create NTHash hash from string password
func GetHash(pass string) string {
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	encoded, _ := utf16.NewEncoder().String(pass)
	h := md4.New()
	io.WriteString(h, encoded)
	return hex.EncodeToString(h.Sum(nil))
}

//CheckHash create NTHash hash from string password
func CheckHash(pass, hash string) (c bool) {
	nthash := GetHash(pass)
	if hash == nthash {
		c = true
	}
	return c
}
