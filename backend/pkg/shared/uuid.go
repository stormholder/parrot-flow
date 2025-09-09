package shared

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateRandomBytes(n int) []byte {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return nil
	}
	return bytes
}

func CustomUUID() string {
	const uuidLength = 16

	version := new(big.Int).SetBytes(generateRandomBytes(4))
	variant := new(big.Int).SetBytes(generateRandomBytes(16))
	node := new(big.Int).SetBytes(generateRandomBytes(8))

	uuidStr := fmt.Sprintf("%02X%02X-%05X-", version, variant, node)
	uuidStr += fmt.Sprintf("%s", uuidStr[:uuidLength-9])
	return uuidStr
}
