package auth

import (
	"crypto/sha512"
	"fmt"
	"jwt-auth/initializators"
)

func Encode(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(initializators.Config.Hash))
}
