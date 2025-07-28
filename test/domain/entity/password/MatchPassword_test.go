package passwordEntityTest

import (
	passwordEntity2 "dev-knowledge/domain/entity/user/password"
	"dev-knowledge/test/domain/entity/entityStub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type MatchPasswordShould struct {
	suite.Suite
}

func TestPasswordMatchShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(MatchPasswordShould))
}

func (s *MatchPasswordShould) TestMatchPassword_PasswordIsEmpty_ReturnError() {
	globalSalt := entityStub.GetPasswordGlobalSalt()

	password, err := passwordEntity2.NewPassword("123password", globalSalt)
	assert.NoError(s.T(), err)
	require.NotNil(s.T(), password)

	err = password.MatchPassword("", globalSalt)
	assert.ErrorIs(s.T(), err, passwordEntity2.ErrPlainTextPasswordIsRequired)
}

func (s *MatchPasswordShould) TestMatchPassword_GlobalSaltIsEmpty_ReturnError() {
	plainTextPassword := "123password"

	password := entityStub.GetPasswordWith(plainTextPassword)

	err := password.MatchPassword(plainTextPassword, "")
	assert.ErrorIs(s.T(), err, passwordEntity2.ErrGlobalSaltIsRequired)
}

func (s *MatchPasswordShould) TestMatchPassword_MismatchedPassword_ReturnError() {
	globalSalt := entityStub.GetPasswordGlobalSalt()

	password, err := passwordEntity2.NewPassword("123password", globalSalt)
	assert.NoError(s.T(), err)
	require.NotNil(s.T(), password)

	err = password.MatchPassword("999", globalSalt)
	assert.ErrorIs(s.T(), err, passwordEntity2.ErrMismatchedPassword)
}

func (s *MatchPasswordShould) TestMatchPassword_MismatchedGlobalSalt_ReturnError() {
	plainTextPassword := "123password"

	password := entityStub.GetPasswordWith(plainTextPassword)

	err := password.MatchPassword(plainTextPassword, entityStub.GetPasswordGlobalSalt())
	assert.ErrorIs(s.T(), err, passwordEntity2.ErrMismatchedPassword)
}

func (s *MatchPasswordShould) TestMatchPassword_EqualPasswords_ReturnNoError() {
	plainTextPassword := "123password"
	globalSalt := entityStub.GetPasswordGlobalSalt()

	password, err := passwordEntity2.NewPassword(plainTextPassword, globalSalt)
	assert.NoError(s.T(), err)
	require.NotNil(s.T(), password)

	err = password.MatchPassword(plainTextPassword, globalSalt)
	assert.NoError(s.T(), err)
}

func (s *MatchPasswordShould) TestMatchPassword_PasswordWithSpaces_ReturnNoError() {
	plainTextPassword := "    123password    "
	globalSalt := entityStub.GetPasswordGlobalSalt()

	password, err := passwordEntity2.NewPassword(plainTextPassword, globalSalt)
	assert.NoError(s.T(), err)
	require.NotNil(s.T(), password)

	err = password.MatchPassword(strings.TrimSpace(plainTextPassword), globalSalt)
	assert.NoError(s.T(), err)
}
