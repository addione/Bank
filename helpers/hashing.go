package helpers

import "golang.org/x/crypto/bcrypt"

type Hashing struct {
}

func newHashing() *Hashing {
	return &Hashing{}
}

func (h *Hashing) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
