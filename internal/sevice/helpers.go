package service

import (
	"crypto/md5"
	"fmt"
)

func (s *userService) hashMd5(in string) string {
	h := md5.New()
	h.Write([]byte(in))

	return fmt.Sprintf("%x", h.Sum(s.salt))
}
