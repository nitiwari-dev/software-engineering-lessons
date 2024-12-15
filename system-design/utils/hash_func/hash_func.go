package hash_func

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashFunction Hash function are functions which takes arbitratry size input and convert them into fixed sized output called as hashcode, and the domain is called hash space
func HashFunction(data string) string {
	if len(data) == 0 {
		return data
	}
	sum := hashInBytes(data)
	encode := hex.EncodeToString(sum)
	return encode
}

func hashInBytes(data string) []byte {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hash.Sum(nil)
}
