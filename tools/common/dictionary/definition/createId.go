package dictionaryDefinition

import (
	"crypto/sha256"
	"fmt"
)

func CreateId(s *string) string {
	hash := sha256.New()
	return fmt.Sprintf("%x", hash.Sum([]byte(*s)))
}
