package email

import (
	"dev-knowledge/test/domain/entity/entityStub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserEmailGenerateCodeShould struct {
	suite.Suite
}

func TestUserEmailGenerateCodeShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserEmailGenerateCodeShould))
}

func (s *UserEmailGenerateCodeShould) TestGenerateCode_ReturnCode() {
	userEmail := entityStub.GetUserEmailWithoutVerify("email@email.com")

	assert.Empty(s.T(), userEmail.VerificationCode())

	err := userEmail.InitVerificationCode()
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), userEmail.VerificationCode())
}
