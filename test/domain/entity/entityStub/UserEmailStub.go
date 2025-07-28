package entityStub

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	emailEntity2 "dev-knowledge/domain/entity/user/email"
)

func GetUserEmailWithoutVerify(emailStr string) *emailEntity2.UserEmail {
	email, err := emailPrimitive.EmailFrom(emailStr)
	if err != nil {
		panic(err)
	}

	userEmail, err := emailEntity2.NewBuilder().
		Email(email).
		Build()

	if err != nil {
		panic(err)
	}

	return userEmail
}

func GetUserEmailWithVerify(emailStr string, verificationCodeStr string) *emailEntity2.UserEmail {
	email, err := emailPrimitive.EmailFrom(emailStr)
	if err != nil {
		panic(err)
	}

	verificationCode, err := verificationPrimitive.NewVerificationCode(verificationCodeStr, 10)
	if err != nil {
		panic(err)
	}

	userEmail, err := emailEntity2.NewBuilder().
		Email(email).
		VerificationCode(verificationCode).
		Build()

	if err != nil {
		panic(err)
	}

	return userEmail
}
