package entityStub

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	emailPrimitiveStub "dev-knowledge/common/domainPrimitive/test/testDouble/stub/email"
	"dev-knowledge/domain/entity/user/email"
)

func GetNonActivatedUserEmails() *emailEntity.UserEmails {
	emailStr := emailPrimitiveStub.GetEmail()
	email, err := emailPrimitive.EmailFrom(string(emailStr))
	if err != nil {
		panic(err)
	}

	userEmails, err := emailEntity.NewEmails(email)
	if err != nil {
		panic(err)
	}

	return userEmails
}

func GetActivatedUserEmails() *emailEntity.UserEmails {
	emailStr := emailPrimitiveStub.GetEmail()
	email, err := emailPrimitive.EmailFrom(string(emailStr))
	if err != nil {
		panic(err)
	}

	userEmails, err := emailEntity.NewEmails(email)
	if err != nil {
		panic(err)
	}

	activationCode, err := userEmails.GenerateActivationCode()
	if err != nil {
		panic(err)
	}

	err = userEmails.Activate(activationCode)
	if err != nil {
		panic(err)
	}

	return userEmails
}
