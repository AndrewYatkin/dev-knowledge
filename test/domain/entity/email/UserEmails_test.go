package email

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	emailEntity2 "dev-knowledge/domain/entity/user/email"
	"dev-knowledge/test/domain/entity/entityStub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserEmailsShould struct {
	suite.Suite
}

func TestUserEmailsShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserEmailsShould))
}

func (s *UserEmailsShould) TestInitNewEmail_ValidEmail_ReturnNonActivated() {
	expectedEmailStr := "email@email.com"
	expectedEmail, err := emailPrimitive.EmailFrom(expectedEmailStr)
	assert.Nil(s.T(), err)

	actualUserEmails, err := emailEntity2.NewEmails(expectedEmail)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualUserEmails)
	assert.EqualValues(s.T(), expectedEmailStr, actualUserEmails.NonActivatedEmail())
	assert.Empty(s.T(), actualUserEmails.ActivatedEmail())

	actualNonActivatedEmail, exists := actualUserEmails.NonActivatedEmailData()
	assert.True(s.T(), exists)
	assert.NotNil(s.T(), actualNonActivatedEmail)

	actualActivatedEmail, exists := actualUserEmails.ActivatedEmailData()
	assert.NotNil(s.T(), actualActivatedEmail)
	assert.Empty(s.T(), actualActivatedEmail.Email())
	assert.False(s.T(), exists)
}

func (s *UserEmailsShould) TestInitNewEmail_NotValidEmail_ReturnError() {
	notValidEmail := emailPrimitive.Email("")

	actualUserEmails, err := emailEntity2.NewEmails(notValidEmail)
	assert.Nil(s.T(), actualUserEmails)
	assert.NotNil(s.T(), err)
}

func (s *UserEmailsShould) TestGenerateActivationCode_NonActivatedNotExist_ReturnError() {
	actualUserEmails := emailEntity2.UserEmails{}

	code, err := actualUserEmails.GenerateActivationCode()
	assert.Empty(s.T(), code)
	assert.Equal(s.T(), emailEntity2.ErrNonActivatedEmailNotExist, err)
}

func (s *UserEmailsShould) TestGenerateActivationCode_NonActivatedExist_ReturnCodeNoErr() {
	emailStr := "email@email.com"
	userEmail, err := emailPrimitive.EmailFrom(emailStr)
	assert.Nil(s.T(), err)

	actualUserEmails, err := emailEntity2.NewEmails(userEmail)
	assert.NotNil(s.T(), actualUserEmails)
	assert.Nil(s.T(), err)

	activationCode, err := actualUserEmails.GenerateActivationCode()
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), activationCode)
}

func (s *UserEmailsShould) TestGenerateActivationCode_TimeoutNotExpired_ReturnError() {
	emailsStub := entityStub.GetNonActivatedUserEmails()
	_, err := emailsStub.GenerateActivationCode()
	assert.Nil(s.T(), err)

	code, err := emailsStub.GenerateActivationCode()
	assert.Empty(s.T(), code)
	assert.Equal(s.T(), emailEntity2.ErrTimeoutHasNotExpired, err)
}

func (s *UserEmailsShould) TestActivate_ExistsNonActivatedEmail_EmailIsActivated() {
	emailStr := "email@email.com"
	userEmail, err := emailPrimitive.EmailFrom(emailStr)
	assert.Nil(s.T(), err)

	actualUserEmails, err := emailEntity2.NewEmails(userEmail)
	assert.NotNil(s.T(), actualUserEmails)
	assert.Nil(s.T(), err)

	activationCode, err := actualUserEmails.GenerateActivationCode()
	assert.Nil(s.T(), err)
	assert.False(s.T(), actualUserEmails.HasActivated())

	err = actualUserEmails.Activate(activationCode)
	assert.Nil(s.T(), err)
	assert.True(s.T(), actualUserEmails.HasActivated())
	assert.Empty(s.T(), actualUserEmails.NonActivatedEmail())
	assert.EqualValues(s.T(), emailStr, actualUserEmails.ActivatedEmail())

	actualNonActivatedEmail, exists := actualUserEmails.NonActivatedEmailData()
	assert.NotNil(s.T(), actualNonActivatedEmail)
	assert.Empty(s.T(), actualNonActivatedEmail.Email())
	assert.False(s.T(), exists)

	actualActivatedEmail, exists := actualUserEmails.ActivatedEmailData()
	assert.True(s.T(), exists)
	assert.NotNil(s.T(), actualActivatedEmail)
}

func (s *UserEmailsShould) TestActivate_NonActivatedEmailNotExist_ReturnError() {
	actualUserEmails := emailEntity2.UserEmails{}

	err := actualUserEmails.Activate("123")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), emailEntity2.ErrNonActivatedEmailNotExist, err)
}

func (s *UserEmailsShould) TestObjInitNewEmail_ValidEmail_ReturnNotActivated() {
	activatedUserEmails := entityStub.GetActivatedUserEmails()
	originalActivatedEmail := activatedUserEmails.ActivatedEmail()

	expectedEmailStr := "email@email.com"
	expectedEmail, err := emailPrimitive.EmailFrom(expectedEmailStr)
	assert.Nil(s.T(), err)

	err = activatedUserEmails.InitNewEmail(expectedEmail)
	assert.Nil(s.T(), err)
	assert.EqualValues(s.T(), originalActivatedEmail, activatedUserEmails.ActivatedEmail())
	assert.EqualValues(s.T(), expectedEmailStr, activatedUserEmails.NonActivatedEmail())
}

func (s *UserEmailsShould) TestObjInitNewEmail_NotValidEmail_ReturnError() {
	activatedUserEmails := entityStub.GetActivatedUserEmails()

	expectedValidEmailStr := "email@email.com"
	expectedValidEmail, err := emailPrimitive.EmailFrom(expectedValidEmailStr)
	assert.Nil(s.T(), err)

	err = activatedUserEmails.InitNewEmail(expectedValidEmail)
	assert.Nil(s.T(), err)
	assert.EqualValues(s.T(), expectedValidEmailStr, activatedUserEmails.NonActivatedEmail())

	notValidEmail := emailPrimitive.Email("")
	err = activatedUserEmails.InitNewEmail(notValidEmail)
	assert.NotNil(s.T(), err)
	assert.EqualValues(s.T(), expectedValidEmailStr, activatedUserEmails.NonActivatedEmail())
}
