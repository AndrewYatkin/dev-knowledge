package passwordEntity

import (
	"dev-knowledge/infrastructure/errors"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost int = 10

type Hash string

func NewHash(plainTextPassword string) (Hash, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcryptCost)
	if err != nil {
		return "", err
	}

	return Hash(hashBytes), nil
}

func HashFrom(hashStr string) (Hash, error) {
	if hashStr == "" {
		return "", nil
	}

	return Hash(hashStr), nil
}

func (h Hash) String() string {
	return string(h)
}

func (h Hash) ComparePassword(plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(plainPassword))
	if err != nil {
		wrapBcryptError(err)
	}

	return nil
}

func wrapBcryptError(err error) error {
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrMismatchedPassword
	}

	return errors.NewError(ComparePasswordUnknownErrorCode, err.Error())
}
