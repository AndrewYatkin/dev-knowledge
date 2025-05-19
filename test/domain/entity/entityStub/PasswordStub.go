package entityStub

import (
	passwordEntity "dev-knowledge/domain/entity/password"
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

func GetPassword() *passwordEntity.Password {
	return GetPasswordWith(GetRandomPlainTextPassword())
}

func GetPasswordWith(plainTextPassword string) *passwordEntity.Password {
	password, err := passwordEntity.NewPassword(plainTextPassword, GetPasswordGlobalSalt())
	if err != nil {
		panic(err)
	}

	return password
}

func GetPasswordBuilder() *passwordEntity.Builder {
	return passwordEntity.NewBuilder().
		Hash(GetPasswordHash()).
		Salt(GetPasswordGlobalSalt())
}

func GetPasswordHash() passwordEntity.Hash {
	hash, err := passwordEntity.NewHash(commonTesting.RandomDefaultStr())
	if err != nil {
		panic(err)
	}

	return hash
}

func GetPasswordGlobalSalt() passwordEntity.Salt {
	return passwordEntity.NewSalt()
}
