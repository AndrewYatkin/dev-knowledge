package entityStub

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	emailEntity "dev-knowledge/domain/entity/email"
)

func GetUserEmailWithoutVerify(emailStr string) *emailEntity.UserEmail {
	email, err := emailPrimitive.EmailFrom(emailStr)
	if err != nil {
		panic(err)
	}

	userEmail, err := emailEntity.NewBuilder().
		Email(email).
		Build()

	if err != nil {
		panic(err)
	}

	return userEmail
}

func GetUserEmailWithVerify(emailStr string, verificationCodeStr string) *emailEntity.UserEmail {
	email, err := emailPrimitive.EmailFrom(emailStr)
	if err != nil {
		panic(err)
	}

	verificationCode, err := verificationPrimitive.NewVerificationCode(verificationCodeStr, 10)
	if err != nil {
		panic(err)
	}

	userEmail, err := emailEntity.NewBuilder().
		Email(email).
		VerificationCode(verificationCode).
		Build()

	if err != nil {
		panic(err)
	}

	return userEmail
}
