package passwordEntity

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	prohibitedSymbolsCheck = regexp.MustCompile("[^a-zA-Z!@#$%^&*(){}\\[\\]\\d_-]+")
	requiredSymbolsCheck   = regexp.MustCompile("([0-9]+.+?[A-z]+)|([A-z].+?[0-9]+)")
)

type Password struct {
	hash Hash
	salt Salt
}

func NewPassword(plainTextPassword string, globalSalt Salt) (*Password, error) {
	if plainTextPassword == "" {
		return nil, ErrPlainTextPasswordIsRequired
	}
	if globalSalt == "" {
		return nil, ErrGlobalSaltIsRequired
	}

	plainTextPassword = strings.TrimSpace(plainTextPassword)
	err := validatePassword(plainTextPassword)
	if err != nil {
		return nil, err
	}

	privateSalt := NewSalt()
	saltedPassword := saltPassword(plainTextPassword, privateSalt, globalSalt)
	passwordHash, err := NewHash(saltedPassword)
	if err != nil {
		return nil, err
	}

	password, err := NewBuilder().
		Hash(passwordHash).
		Salt(privateSalt).
		Build()
	if err != nil {
		return nil, err
	}

	return password, nil
}

func validatePassword(plainTextPassword string) error {
	passRune := []rune(plainTextPassword)
	if len(passRune) < 6 {
		return ErrShortPassword
	}
	if len(passRune) > 50 {
		return ErrTooLongPassword
	}

	containProhibitedSymbols := prohibitedSymbolsCheck.MatchString(plainTextPassword)
	if containProhibitedSymbols {
		return ErrPasswordContainsProhibitedSymbols
	}

	containAllRequiredSymbols := requiredSymbolsCheck.MatchString(plainTextPassword)
	if !containAllRequiredSymbols {
		return ErrPasswordNotContainRequiredSymbols
	}

	return nil
}

func (p *Password) MatchPassword(plainTextPassword string, globalSalt Salt) error {
	if plainTextPassword == "" {
		return ErrPlainTextPasswordIsRequired
	}
	if globalSalt == "" {
		return ErrGlobalSaltIsRequired
	}

	saltedPassword := saltPassword(plainTextPassword, p.salt, globalSalt)
	err := p.hash.ComparePassword(saltedPassword)
	if err != nil {
		return err
	}

	return nil
}

func saltPassword(plainTextPassword string, privateSalt, globalSalt Salt) string {
	return fmt.Sprintf("%s%s%s", globalSalt, plainTextPassword, privateSalt)
}
