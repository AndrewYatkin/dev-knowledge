package entityStub

import (
	passwordEntity2 "dev-knowledge/domain/entity/user/password"
	commonTesting "dev-knowledge/infrastructure/testing"
	"fmt"
)

const (
	minRandomPasswordLength = 100000
	maxRandomPasswordLength = 999999
)

func GetRandomPlainTextPassword() string {
	return fmt.Sprintf("abc%d", commonTesting.RandomNumber(minRandomPasswordLength, maxRandomPasswordLength))
}

func GetPassword() *passwordEntity2.Password {
	return GetPasswordWith(GetRandomPlainTextPassword())
}

func GetPasswordWith(plainTextPassword string) *passwordEntity2.Password {
	password, err := passwordEntity2.NewPassword(plainTextPassword, GetPasswordGlobalSalt())
	if err != nil {
		panic(err)
	}

	return password
}

func GetPasswordBuilder() *passwordEntity2.Builder {
	return passwordEntity2.NewBuilder().
		Hash(GetPasswordHash()).
		Salt(GetPasswordGlobalSalt())
}

func GetPasswordHash() passwordEntity2.Hash {
	hash, err := passwordEntity2.NewHash(commonTesting.RandomDefaultStr())
	if err != nil {
		panic(err)
	}

	return hash
}

func GetPasswordGlobalSalt() passwordEntity2.Salt {
	return passwordEntity2.NewSalt()
}
