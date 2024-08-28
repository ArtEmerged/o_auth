package user

import (
	"crypto/sha256"
	"fmt"
)

func (s *userService) hashSha256(in string) string {
	h := sha256.New()

	h.Write(s.salt)
	h.Write([]byte(in))

	return fmt.Sprintf("%x", h.Sum(nil))
}
