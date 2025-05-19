package passwordEntityTest

import (
	passwordEntity "dev-knowledge/domain/entity/password"
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

	hash, err := passwordEntity.NewHash(plainTextPassword)

	assert.NotEmpty(s.T(), hash)
	assert.NoError(s.T(), err)
}

func (s *HashShould) TestHashFrom_EmptyHash_ReturnError() {
	hash, err := passwordEntity.HashFrom("")

	assert.Empty(s.T(), hash)
	assert.ErrorIs(s.T(), err, passwordEntity.ErrHashIsEmpty)
}

func (s *HashShould) TestHashFrom_ValidHash_ReturnHash() {
	expectedHashStr := "12345"

	hash, err := passwordEntity.HashFrom(expectedHashStr)

	assert.EqualValues(s.T(), expectedHashStr, hash)
	assert.NoError(s.T(), err)
}

func (s *HashShould) TestComparePassword_PasswordMismatch_ReturnError() {
	originalPlainTextPassword := "12345"

	hash, err := passwordEntity.NewHash(originalPlainTextPassword)
	assert.NoError(s.T(), err)
	require.NotEmpty(s.T(), hash)

	err = hash.ComparePassword("54321")
	assert.ErrorIs(s.T(), err, passwordEntity.ErrMismatchedPassword)
}

func (s *HashShould) TestComparePassword_PasswordEquals_ReturnNoError() {
	plainTextPassword := "12345"

	hash, err := passwordEntity.NewHash(plainTextPassword)
	assert.NoError(s.T(), err)
	require.NotEmpty(s.T(), hash)

	err = hash.ComparePassword(plainTextPassword)
	assert.NoError(s.T(), err)
}
