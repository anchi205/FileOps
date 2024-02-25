package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// get []byte as input and return the SHA1 hash as a string
func GetSHA1(data []byte) (string, error) {
	h := sha1.New()

	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	// Get the final hash and convert it to a hexadecimal string
	hashInBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}
