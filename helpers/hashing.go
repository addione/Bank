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

func (h *Hashing) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
