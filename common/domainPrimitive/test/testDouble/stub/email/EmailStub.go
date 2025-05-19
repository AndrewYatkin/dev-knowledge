package emailPrimitiveStub

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	commonTesting "dev-knowledge/infrastructure/testing"
	"fmt"
)

func GetEmail() emailPrimitive.Email {
	login := commonTesting.RandomDefaultStr()
	mail := commonTesting.RandomStr(7)
	emailStr := fmt.Sprintf("%s@%s.test", login, mail)

	email, err := emailPrimitive.EmailFrom(emailStr)
	if err != nil {
		panic(err)
	}

	return email
}
