package verification

import (
	commonTime "dev-knowledge/infrastructure/tools/time"
	"time"
)

func GetVerificationCodeStub(code string) *verificationEntity.VerificationCode {
	expireAt := commonTime.Now().Add(time.Minute * time.Duration(10))

	return GetVerificationCodeStubByExpireAt(code, expireAt)
}

func GetVerificationCodeStubByExpireAt(code string, expireAt *commonTime.Time) *verificationEntity.VerificationCode {
	verificationCode, err := verificationEntity.NewBuilder().
		Code(code).
		CreatedAt(expireAt.Add(time.Minute * time.Duration(-10))).
		ExpireAt(expireAt).
		Build()

	if err != nil {
		panic(err)
	}

	return verificationCode
}
