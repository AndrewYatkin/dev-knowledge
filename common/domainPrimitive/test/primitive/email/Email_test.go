package email

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EmailShould struct {
	suite.Suite
}

func TestEmailShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(EmailShould))
}

func (s *EmailShould) TestEmailFrom_ValidEmail_ReturnEmail() {
	validEmailCases := []struct {
		sourceEmail   string
		expectedEmail string
	}{
		{"test@mail.ru", "test@mail.ru"},
		{"123@mail.com", "123@mail.com"},
		{"MY123@google.com", "my123@google.com"},
		{"mY.321@ya.com", "my.321@ya.com"},
		{"my.321-email@EMail.com", "my.321-email@email.com"},
		{"test_email@mail.com", "test_email@mail.com"},
		{"test@mail123.com", "test@mail123.com"},
		{"test@mail.com.ru", "test@mail.com.ru"},
		{"  teST@mail.com.ru  ", "test@mail.com.ru"},
	}

	for _, testCase := range validEmailCases {
		aclualEmail, err := emailPrimitive.EmailFrom(testCase.sourceEmail)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), testCase.expectedEmail, aclualEmail.String())
	}
}

func (s *EmailShould) TestEmailFrom_EmptyEmail_ReturnError() {
	notValidEmailValues := []string{
		"",
		"    ",
	}

	for _, notValidEmailValue := range notValidEmailValues {
		_, err := emailPrimitive.EmailFrom(notValidEmailValue)

		assert.Equal(s.T(), emailPrimitive.ErrWrongEmail, err)
	}

}

func (s *EmailShould) TestEmailFrom_NotValidEmail_ReturnError() {
	notValidEmailsValues := []string{
		"@mail.ru",
		"testmail.ru",
		"test@mail",
		"  test@mail  ",
	}

	for _, notValidEmailValue := range notValidEmailsValues {
		_, err := emailPrimitive.EmailFrom(notValidEmailValue)

		assert.Equalf(s.T(), emailPrimitive.ErrWrongEmail, err, "not assert email = %s", notValidEmailValue)
	}
}
