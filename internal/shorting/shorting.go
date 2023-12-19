package shorting

import (
	"math/rand"
	"time"
)

const (
	// Length of Identifier
	Length = 10

	// CharSet used to create identifier
	CharSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
)

// Generate a random identifier
func GenerateIdentifier() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random string from the specified charset
	randomString := ""
	for i := 0; i < Length; i++ {
		randomString += string(CharSet[r.Intn(len(CharSet))])
	}

	return randomString
}
