package passwordEntityTest

import (
	passwordEntity2 "dev-knowledge/domain/entity/user/password"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HashShould struct {
	suite.Suite
}

func TestHashShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(HashShould))
}

func (s *HashShould) TestNewHash_ValidParams_ReturnHash() {
	plainTextPassword := "12345"

	hash, err := passwordEntity2.NewHash(plainTextPassword)

	assert.NotEmpty(s.T(), hash)
	assert.NoError(s.T(), err)
}

func (s *HashShould) TestHashFrom_EmptyHash_ReturnError() {
	hash, err := passwordEntity2.HashFrom("")

	assert.Empty(s.T(), hash)
	assert.ErrorIs(s.T(), err, passwordEntity2.ErrHashIsEmpty)
}

func (s *HashShould) TestHashFrom_ValidHash_ReturnHash() {
	expectedHashStr := "12345"

	hash, err := passwordEntity2.HashFrom(expectedHashStr)

	assert.EqualValues(s.T(), expectedHashStr, hash)
	assert.NoError(s.T(), err)
}

func (s *HashShould) TestComparePassword_PasswordMismatch_ReturnError() {
	originalPlainTextPassword := "12345"

	hash, err := passwordEntity2.NewHash(originalPlainTextPassword)
	assert.NoError(s.T(), err)
	require.NotEmpty(s.T(), hash)

	err = hash.ComparePassword("54321")
	assert.ErrorIs(s.T(), err, passwordEntity2.ErrMismatchedPassword)
}

func (s *HashShould) TestComparePassword_PasswordEquals_ReturnNoError() {
	plainTextPassword := "12345"

	hash, err := passwordEntity2.NewHash(plainTextPassword)
	assert.NoError(s.T(), err)
	require.NotEmpty(s.T(), hash)

	err = hash.ComparePassword(plainTextPassword)
	assert.NoError(s.T(), err)
}
