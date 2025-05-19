package entityStub

import (
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	"math/rand"
	"strconv"
)

func GetVerificationCode() *verificationPrimitive.VerificationCode {
	randomCode := strconv.Itoa(1000 + rand.Intn(9000))

	verificationCode, err := verificationPrimitive.NewVerificationCode(randomCode, 10)
	if err != nil {
		panic(err)
	}

	return verificationCode
}
