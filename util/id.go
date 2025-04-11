package util

import (
	"crypto/rand"
	"fmt"
)

func GenId(prefix string) string {
	randomBytes := make([]byte, 12)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s_%x", prefix, randomBytes)
}
