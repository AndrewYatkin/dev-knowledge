package passwordEntity

import "github.com/google/uuid"

const (
	minSaltLength = 5
	maxSaltLength = 10
)

type Salt string

func NewSalt() Salt {
	salt := uuid.NewString()
	if len(salt) > maxSaltLength {
		salt = salt[:maxSaltLength]
	}
	return Salt(salt)
}

func SaltFrom(salt string) (Salt, error) {
	if len(salt) < minSaltLength {
		return "", ErrShortSalt
	}
	if len(salt) > maxSaltLength {
		return "", ErrLongSalt
	}
	return Salt(salt), nil
}

func (s Salt) String() string {
	return string(s)
}
